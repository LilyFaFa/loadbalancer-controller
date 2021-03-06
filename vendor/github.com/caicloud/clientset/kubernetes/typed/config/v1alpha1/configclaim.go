/*
Copyright 2017 caicloud authors. All rights reserved.
*/

package v1alpha1

import (
	scheme "github.com/caicloud/clientset/kubernetes/scheme"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ConfigClaimsGetter has a method to return a ConfigClaimInterface.
// A group's client should implement this interface.
type ConfigClaimsGetter interface {
	ConfigClaims(namespace string) ConfigClaimInterface
}

// ConfigClaimInterface has methods to work with ConfigClaim resources.
type ConfigClaimInterface interface {
	Create(*v1alpha1.ConfigClaim) (*v1alpha1.ConfigClaim, error)
	Update(*v1alpha1.ConfigClaim) (*v1alpha1.ConfigClaim, error)
	UpdateStatus(*v1alpha1.ConfigClaim) (*v1alpha1.ConfigClaim, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ConfigClaim, error)
	List(opts v1.ListOptions) (*v1alpha1.ConfigClaimList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigClaim, err error)
	ConfigClaimExpansion
}

// configClaims implements ConfigClaimInterface
type configClaims struct {
	client rest.Interface
	ns     string
}

// newConfigClaims returns a ConfigClaims
func newConfigClaims(c *ConfigV1alpha1Client, namespace string) *configClaims {
	return &configClaims{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a configClaim and creates it.  Returns the server's representation of the configClaim, and an error, if there is any.
func (c *configClaims) Create(configClaim *v1alpha1.ConfigClaim) (result *v1alpha1.ConfigClaim, err error) {
	result = &v1alpha1.ConfigClaim{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configclaims").
		Body(configClaim).
		Do().
		Into(result)
	return
}

// Update takes the representation of a configClaim and updates it. Returns the server's representation of the configClaim, and an error, if there is any.
func (c *configClaims) Update(configClaim *v1alpha1.ConfigClaim) (result *v1alpha1.ConfigClaim, err error) {
	result = &v1alpha1.ConfigClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configclaims").
		Name(configClaim.Name).
		Body(configClaim).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclientstatus=false comment above the type to avoid generating UpdateStatus().

func (c *configClaims) UpdateStatus(configClaim *v1alpha1.ConfigClaim) (result *v1alpha1.ConfigClaim, err error) {
	result = &v1alpha1.ConfigClaim{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configclaims").
		Name(configClaim.Name).
		SubResource("status").
		Body(configClaim).
		Do().
		Into(result)
	return
}

// Delete takes name of the configClaim and deletes it. Returns an error if one occurs.
func (c *configClaims) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configclaims").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configClaims) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configclaims").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the configClaim, and returns the corresponding configClaim object, and an error if there is any.
func (c *configClaims) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigClaim, err error) {
	result = &v1alpha1.ConfigClaim{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configclaims").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConfigClaims that match those selectors.
func (c *configClaims) List(opts v1.ListOptions) (result *v1alpha1.ConfigClaimList, err error) {
	result = &v1alpha1.ConfigClaimList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configClaims.
func (c *configClaims) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configclaims").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched configClaim.
func (c *configClaims) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigClaim, err error) {
	result = &v1alpha1.ConfigClaim{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configclaims").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
