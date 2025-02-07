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
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/kubedb/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MySQLsGetter has a method to return a MySQLInterface.
// A group's client should implement this interface.
type MySQLsGetter interface {
	MySQLs(namespace string) MySQLInterface
}

// MySQLInterface has methods to work with MySQL resources.
type MySQLInterface interface {
	Create(*v1alpha1.MySQL) (*v1alpha1.MySQL, error)
	Update(*v1alpha1.MySQL) (*v1alpha1.MySQL, error)
	UpdateStatus(*v1alpha1.MySQL) (*v1alpha1.MySQL, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MySQL, error)
	List(opts v1.ListOptions) (*v1alpha1.MySQLList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MySQL, err error)
	MySQLExpansion
}

// mySQLs implements MySQLInterface
type mySQLs struct {
	client rest.Interface
	ns     string
}

// newMySQLs returns a MySQLs
func newMySQLs(c *KubedbV1alpha1Client, namespace string) *mySQLs {
	return &mySQLs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mySQL, and returns the corresponding mySQL object, and an error if there is any.
func (c *mySQLs) Get(name string, options v1.GetOptions) (result *v1alpha1.MySQL, err error) {
	result = &v1alpha1.MySQL{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MySQLs that match those selectors.
func (c *mySQLs) List(opts v1.ListOptions) (result *v1alpha1.MySQLList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MySQLList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mySQLs.
func (c *mySQLs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mySQL and creates it.  Returns the server's representation of the mySQL, and an error, if there is any.
func (c *mySQLs) Create(mySQL *v1alpha1.MySQL) (result *v1alpha1.MySQL, err error) {
	result = &v1alpha1.MySQL{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mysqls").
		Body(mySQL).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mySQL and updates it. Returns the server's representation of the mySQL, and an error, if there is any.
func (c *mySQLs) Update(mySQL *v1alpha1.MySQL) (result *v1alpha1.MySQL, err error) {
	result = &v1alpha1.MySQL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqls").
		Name(mySQL.Name).
		Body(mySQL).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *mySQLs) UpdateStatus(mySQL *v1alpha1.MySQL) (result *v1alpha1.MySQL, err error) {
	result = &v1alpha1.MySQL{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mysqls").
		Name(mySQL.Name).
		SubResource("status").
		Body(mySQL).
		Do().
		Into(result)
	return
}

// Delete takes name of the mySQL and deletes it. Returns an error if one occurs.
func (c *mySQLs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqls").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mySQLs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mysqls").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mySQL.
func (c *mySQLs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MySQL, err error) {
	result = &v1alpha1.MySQL{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mysqls").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
