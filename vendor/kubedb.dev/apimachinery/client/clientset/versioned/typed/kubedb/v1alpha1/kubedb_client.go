/*
Copyright 2019 The KubeDB Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	"kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type KubedbV1alpha1Interface interface {
	RESTClient() rest.Interface
	DormantDatabasesGetter
	ElasticsearchesGetter
	EtcdsGetter
	MariaDBsGetter
	MemcachedsGetter
	MongoDBsGetter
	MySQLsGetter
	PerconaXtraDBsGetter
	PgBouncersGetter
	PostgresesGetter
	ProxySQLsGetter
	RedisesGetter
	SnapshotsGetter
}

// KubedbV1alpha1Client is used to interact with features provided by the kubedb.com group.
type KubedbV1alpha1Client struct {
	restClient rest.Interface
}

func (c *KubedbV1alpha1Client) DormantDatabases(namespace string) DormantDatabaseInterface {
	return newDormantDatabases(c, namespace)
}

func (c *KubedbV1alpha1Client) Elasticsearches(namespace string) ElasticsearchInterface {
	return newElasticsearches(c, namespace)
}

func (c *KubedbV1alpha1Client) Etcds(namespace string) EtcdInterface {
	return newEtcds(c, namespace)
}

func (c *KubedbV1alpha1Client) MariaDBs(namespace string) MariaDBInterface {
	return newMariaDBs(c, namespace)
}

func (c *KubedbV1alpha1Client) Memcacheds(namespace string) MemcachedInterface {
	return newMemcacheds(c, namespace)
}

func (c *KubedbV1alpha1Client) MongoDBs(namespace string) MongoDBInterface {
	return newMongoDBs(c, namespace)
}

func (c *KubedbV1alpha1Client) MySQLs(namespace string) MySQLInterface {
	return newMySQLs(c, namespace)
}

func (c *KubedbV1alpha1Client) PerconaXtraDBs(namespace string) PerconaXtraDBInterface {
	return newPerconaXtraDBs(c, namespace)
}

func (c *KubedbV1alpha1Client) PgBouncers(namespace string) PgBouncerInterface {
	return newPgBouncers(c, namespace)
}

func (c *KubedbV1alpha1Client) Postgreses(namespace string) PostgresInterface {
	return newPostgreses(c, namespace)
}

func (c *KubedbV1alpha1Client) ProxySQLs(namespace string) ProxySQLInterface {
	return newProxySQLs(c, namespace)
}

func (c *KubedbV1alpha1Client) Redises(namespace string) RedisInterface {
	return newRedises(c, namespace)
}

func (c *KubedbV1alpha1Client) Snapshots(namespace string) SnapshotInterface {
	return newSnapshots(c, namespace)
}

// NewForConfig creates a new KubedbV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*KubedbV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KubedbV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new KubedbV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubedbV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubedbV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *KubedbV1alpha1Client {
	return &KubedbV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubedbV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
