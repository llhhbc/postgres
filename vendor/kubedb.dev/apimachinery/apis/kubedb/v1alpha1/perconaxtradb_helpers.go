package v1alpha1

import (
	"fmt"

	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/apis/kubedb"

	"github.com/appscode/go/types"
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	crdutils "kmodules.xyz/client-go/apiextensions/v1beta1"
	meta_util "kmodules.xyz/client-go/meta"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	mona "kmodules.xyz/monitoring-agent-api/api/v1"
)

var _ apis.ResourceInfo = &PerconaXtraDB{}

func (p PerconaXtraDB) OffshootName() string {
	return p.Name
}

func (p PerconaXtraDB) OffshootSelectors() map[string]string {
	return map[string]string{
		LabelDatabaseName: p.Name,
		LabelDatabaseKind: ResourceKindPerconaXtraDB,
	}
}

func (p PerconaXtraDB) OffshootLabels() map[string]string {
	out := p.OffshootSelectors()
	out[meta_util.NameLabelKey] = ResourceSingularPerconaXtraDB
	out[meta_util.VersionLabelKey] = string(p.Spec.Version)
	out[meta_util.InstanceLabelKey] = p.Name
	out[meta_util.ComponentLabelKey] = ComponentDatabase
	out[meta_util.ManagedByLabelKey] = GenericKey
	return meta_util.FilterKeys(GenericKey, out, p.Labels)
}

func (p PerconaXtraDB) ResourceShortCode() string {
	return ResourceCodePerconaXtraDB
}

func (p PerconaXtraDB) ResourceKind() string {
	return ResourceKindPerconaXtraDB
}

func (p PerconaXtraDB) ResourceSingular() string {
	return ResourceSingularPerconaXtraDB
}

func (p PerconaXtraDB) ResourcePlural() string {
	return ResourcePluralPerconaXtraDB
}

func (p PerconaXtraDB) ServiceName() string {
	return p.OffshootName()
}

func (p PerconaXtraDB) IsCluster() bool {
	return types.Int32(p.Spec.Replicas) > 1
}

func (p PerconaXtraDB) GoverningServiceName() string {
	return p.OffshootName() + "-gvr"
}

func (p PerconaXtraDB) PeerName(idx int) string {
	return fmt.Sprintf("%s-%d.%s.%s", p.OffshootName(), idx, p.GoverningServiceName(), p.Namespace)
}

func (p PerconaXtraDB) GetDatabaseSecretName() string {
	return p.Spec.DatabaseSecret.SecretName
}

func (p PerconaXtraDB) ClusterName() string {
	return p.OffshootName()
}

type perconaXtraDBApp struct {
	*PerconaXtraDB
}

func (p perconaXtraDBApp) Name() string {
	return p.PerconaXtraDB.Name
}

func (p perconaXtraDBApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularPerconaXtraDB))
}

func (p PerconaXtraDB) AppBindingMeta() appcat.AppBindingMeta {
	return &perconaXtraDBApp{&p}
}

type perconaXtraDBStatsService struct {
	*PerconaXtraDB
}

func (p perconaXtraDBStatsService) GetNamespace() string {
	return p.PerconaXtraDB.GetNamespace()
}

func (p perconaXtraDBStatsService) ServiceName() string {
	return p.OffshootName() + "-stats"
}

func (p perconaXtraDBStatsService) ServiceMonitorName() string {
	return fmt.Sprintf("kubedb-%s-%s", p.Namespace, p.Name)
}

func (p perconaXtraDBStatsService) Path() string {
	return DefaultStatsPath
}

func (p perconaXtraDBStatsService) Scheme() string {
	return ""
}

func (p PerconaXtraDB) StatsService() mona.StatsAccessor {
	return &perconaXtraDBStatsService{&p}
}

func (p PerconaXtraDB) StatsServiceLabels() map[string]string {
	lbl := meta_util.FilterKeys(GenericKey, p.OffshootSelectors(), p.Labels)
	lbl[LabelRole] = RoleStats
	return lbl
}

func (p *PerconaXtraDB) GetMonitoringVendor() string {
	if p.Spec.Monitor != nil {
		return p.Spec.Monitor.Agent.Vendor()
	}
	return ""
}

func (p PerconaXtraDB) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourcePluralPerconaXtraDB,
		Singular:      ResourceSingularPerconaXtraDB,
		Kind:          ResourceKindPerconaXtraDB,
		ShortNames:    []string{ResourceCodePerconaXtraDB},
		Categories:    []string{"datastore", "kubedb", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "kubedb"},
		},
		SpecDefinitionName:      "kubedb.dev/apimachinery/apis/kubedb/v1alpha1.PerconaXtraDB",
		EnableValidation:        true,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: true,
		AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
			{
				Name:     "Version",
				Type:     "string",
				JSONPath: ".spec.version",
			},
			{
				Name:     "Status",
				Type:     "string",
				JSONPath: ".status.phase",
			},
			{
				Name:     "Age",
				Type:     "date",
				JSONPath: ".metadata.creationTimestamp",
			},
		},
	}, apis.SetNameSchema)
}

func (p *PerconaXtraDB) SetDefaults() {
	if p == nil {
		return
	}
	p.Spec.SetDefaults()
}

func (p *PerconaXtraDBSpec) SetDefaults() {
	if p == nil {
		return
	}

	if p.Replicas == nil {
		p.Replicas = types.Int32P(1)
	}

	if p.StorageType == "" {
		p.StorageType = StorageTypeDurable
	}
	if p.UpdateStrategy.Type == "" {
		p.UpdateStrategy.Type = apps.RollingUpdateStatefulSetStrategyType
	}
	if p.TerminationPolicy == "" {
		p.TerminationPolicy = TerminationPolicyDelete
	}
	p.setDefaultProbes()
}

// setDefaultProbes sets defaults only when probe fields are nil.
// In operator, check if the value of probe fields is "{}".
// For "{}", ignore readinessprobe or livenessprobe in statefulset.
// Ref: https://github.com/mattlord/Docker-InnoDB-Cluster/blob/master/healthcheck.sh#L10
func (p *PerconaXtraDBSpec) setDefaultProbes() {
	if p == nil {
		return
	}

	var readynessProbeCmd []string
	if types.Int32(p.Replicas) > 1 {
		readynessProbeCmd = []string{
			"/cluster-check.sh",
		}
	} else {
		readynessProbeCmd = []string{
			"bash",
			"-c",
			`export MYSQL_PWD="${MYSQL_ROOT_PASSWORD}"
ping_resp=$(mysqladmin -uroot ping)
if [[ "$ping_resp" != "mysqld is alive" ]]; then
    echo "[ERROR] server is not ready. PING_RESPONSE: $ping_resp"
    exit 1
fi
`,
		}
	}

	readinessProbe := &core.Probe{
		Handler: core.Handler{
			Exec: &core.ExecAction{
				Command: readynessProbeCmd,
			},
		},
		InitialDelaySeconds: 30,
		PeriodSeconds:       10,
	}
	if p.PodTemplate.Spec.ReadinessProbe == nil {
		p.PodTemplate.Spec.ReadinessProbe = readinessProbe
	}
}

func (p *PerconaXtraDBSpec) GetSecrets() []string {
	if p == nil {
		return nil
	}

	var secrets []string
	if p.DatabaseSecret != nil {
		secrets = append(secrets, p.DatabaseSecret.SecretName)
	}
	return secrets
}
