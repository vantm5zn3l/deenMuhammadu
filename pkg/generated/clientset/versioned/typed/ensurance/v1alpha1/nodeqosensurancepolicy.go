// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/gocrane-io/api/ensurance/v1alpha1"
	scheme "github.com/gocrane-io/api/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeQOSEnsurancePoliciesGetter has a method to return a NodeQOSEnsurancePolicyInterface.
// A group's client should implement this interface.
type NodeQOSEnsurancePoliciesGetter interface {
	NodeQOSEnsurancePolicies(namespace string) NodeQOSEnsurancePolicyInterface
}

// NodeQOSEnsurancePolicyInterface has methods to work with NodeQOSEnsurancePolicy resources.
type NodeQOSEnsurancePolicyInterface interface {
	Create(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.CreateOptions) (*v1alpha1.NodeQOSEnsurancePolicy, error)
	Update(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.UpdateOptions) (*v1alpha1.NodeQOSEnsurancePolicy, error)
	UpdateStatus(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.UpdateOptions) (*v1alpha1.NodeQOSEnsurancePolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NodeQOSEnsurancePolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeQOSEnsurancePolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeQOSEnsurancePolicy, err error)
	NodeQOSEnsurancePolicyExpansion
}

// nodeQOSEnsurancePolicies implements NodeQOSEnsurancePolicyInterface
type nodeQOSEnsurancePolicies struct {
	client rest.Interface
	ns     string
}

// newNodeQOSEnsurancePolicies returns a NodeQOSEnsurancePolicies
func newNodeQOSEnsurancePolicies(c *EnsuranceV1alpha1Client, namespace string) *nodeQOSEnsurancePolicies {
	return &nodeQOSEnsurancePolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeQOSEnsurancePolicy, and returns the corresponding nodeQOSEnsurancePolicy object, and an error if there is any.
func (c *nodeQOSEnsurancePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeQOSEnsurancePolicy, err error) {
	result = &v1alpha1.NodeQOSEnsurancePolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeQOSEnsurancePolicies that match those selectors.
func (c *nodeQOSEnsurancePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeQOSEnsurancePolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeQOSEnsurancePolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeQOSEnsurancePolicies.
func (c *nodeQOSEnsurancePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeQOSEnsurancePolicy and creates it.  Returns the server's representation of the nodeQOSEnsurancePolicy, and an error, if there is any.
func (c *nodeQOSEnsurancePolicies) Create(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.CreateOptions) (result *v1alpha1.NodeQOSEnsurancePolicy, err error) {
	result = &v1alpha1.NodeQOSEnsurancePolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOSEnsurancePolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeQOSEnsurancePolicy and updates it. Returns the server's representation of the nodeQOSEnsurancePolicy, and an error, if there is any.
func (c *nodeQOSEnsurancePolicies) Update(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.UpdateOptions) (result *v1alpha1.NodeQOSEnsurancePolicy, err error) {
	result = &v1alpha1.NodeQOSEnsurancePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		Name(nodeQOSEnsurancePolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOSEnsurancePolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodeQOSEnsurancePolicies) UpdateStatus(ctx context.Context, nodeQOSEnsurancePolicy *v1alpha1.NodeQOSEnsurancePolicy, opts v1.UpdateOptions) (result *v1alpha1.NodeQOSEnsurancePolicy, err error) {
	result = &v1alpha1.NodeQOSEnsurancePolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		Name(nodeQOSEnsurancePolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeQOSEnsurancePolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeQOSEnsurancePolicy and deletes it. Returns an error if one occurs.
func (c *nodeQOSEnsurancePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeQOSEnsurancePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeQOSEnsurancePolicy.
func (c *nodeQOSEnsurancePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeQOSEnsurancePolicy, err error) {
	result = &v1alpha1.NodeQOSEnsurancePolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodeqosensurancepolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
