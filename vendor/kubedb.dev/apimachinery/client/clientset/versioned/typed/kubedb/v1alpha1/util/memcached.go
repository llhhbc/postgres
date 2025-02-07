package util

import (
	"fmt"

	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	cs "kubedb.dev/apimachinery/client/clientset/versioned/typed/kubedb/v1alpha1"

	"github.com/golang/glog"
	kerr "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/jsonmergepatch"
	"k8s.io/apimachinery/pkg/util/wait"
	kutil "kmodules.xyz/client-go"
)

func CreateOrPatchMemcached(c cs.KubedbV1alpha1Interface, meta metav1.ObjectMeta, transform func(*api.Memcached) *api.Memcached) (*api.Memcached, kutil.VerbType, error) {
	cur, err := c.Memcacheds(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
	if kerr.IsNotFound(err) {
		glog.V(3).Infof("Creating Memcached %s/%s.", meta.Namespace, meta.Name)
		out, err := c.Memcacheds(meta.Namespace).Create(transform(&api.Memcached{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Memcached",
				APIVersion: api.SchemeGroupVersion.String(),
			},
			ObjectMeta: meta,
		}))
		return out, kutil.VerbCreated, err
	} else if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	return PatchMemcached(c, cur, transform)
}

func PatchMemcached(c cs.KubedbV1alpha1Interface, cur *api.Memcached, transform func(*api.Memcached) *api.Memcached) (*api.Memcached, kutil.VerbType, error) {
	return PatchMemcachedObject(c, cur, transform(cur.DeepCopy()))
}

func PatchMemcachedObject(c cs.KubedbV1alpha1Interface, cur, mod *api.Memcached) (*api.Memcached, kutil.VerbType, error) {
	curJson, err := json.Marshal(cur)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	modJson, err := json.Marshal(mod)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}

	patch, err := jsonmergepatch.CreateThreeWayJSONMergePatch(curJson, modJson, curJson)
	if err != nil {
		return nil, kutil.VerbUnchanged, err
	}
	if len(patch) == 0 || string(patch) == "{}" {
		return cur, kutil.VerbUnchanged, nil
	}
	glog.V(3).Infof("Patching Memcached %s/%s with %s.", cur.Namespace, cur.Name, string(patch))
	out, err := c.Memcacheds(cur.Namespace).Patch(cur.Name, types.MergePatchType, patch)
	return out, kutil.VerbPatched, err
}

func TryUpdateMemcached(c cs.KubedbV1alpha1Interface, meta metav1.ObjectMeta, transform func(*api.Memcached) *api.Memcached) (result *api.Memcached, err error) {
	attempt := 0
	err = wait.PollImmediate(kutil.RetryInterval, kutil.RetryTimeout, func() (bool, error) {
		attempt++
		cur, e2 := c.Memcacheds(meta.Namespace).Get(meta.Name, metav1.GetOptions{})
		if kerr.IsNotFound(e2) {
			return false, e2
		} else if e2 == nil {
			result, e2 = c.Memcacheds(cur.Namespace).Update(transform(cur.DeepCopy()))
			return e2 == nil, nil
		}
		glog.Errorf("Attempt %d failed to update Memcached %s/%s due to %v.", attempt, cur.Namespace, cur.Name, e2)
		return false, nil
	})

	if err != nil {
		err = fmt.Errorf("failed to update Memcached %s/%s after %d attempts due to %v", meta.Namespace, meta.Name, attempt, err)
	}
	return
}

func UpdateMemcachedStatus(
	c cs.KubedbV1alpha1Interface,
	in *api.Memcached,
	transform func(*api.MemcachedStatus) *api.MemcachedStatus,
) (result *api.Memcached, err error) {
	apply := func(x *api.Memcached) *api.Memcached {
		return &api.Memcached{
			TypeMeta:   x.TypeMeta,
			ObjectMeta: x.ObjectMeta,
			Spec:       x.Spec,
			Status:     *transform(in.Status.DeepCopy()),
		}
	}

	attempt := 0
	cur := in.DeepCopy()
	err = wait.PollImmediate(kutil.RetryInterval, kutil.RetryTimeout, func() (bool, error) {
		attempt++
		var e2 error
		result, e2 = c.Memcacheds(in.Namespace).UpdateStatus(apply(cur))
		if kerr.IsConflict(e2) {
			latest, e3 := c.Memcacheds(in.Namespace).Get(in.Name, metav1.GetOptions{})
			switch {
			case e3 == nil:
				cur = latest
				return false, nil
			case kutil.IsRequestRetryable(e3):
				return false, nil
			default:
				return false, e3
			}
		} else if err != nil && !kutil.IsRequestRetryable(e2) {
			return false, e2
		}
		return e2 == nil, nil
	})

	if err != nil {
		err = fmt.Errorf("failed to update status of Memcached %s/%s after %d attempts due to %v", in.Namespace, in.Name, attempt, err)
	}
	return
}
