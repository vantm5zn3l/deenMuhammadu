// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane-io/api/ensurance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodQOSEnsurancePolicies implements PodQOSEnsurancePolicyInterface
type FakePodQOSEnsurancePolicies struct {
	Fake *FakeEnsuranceV1alpha1
	ns   string
}

var podqosensurancepoliciesResource = schema.GroupVersionResource{Group: "ensurance.crane.io", Version: "v1alpha1", Resource: "podqosensurancepolicies"}

var podqosensurancepoliciesKind = schema.GroupVersionKind{Group: "ensurance.crane.io", Version: "v1alpha1", Kind: "PodQOSEnsurancePolicy"}

// Get takes name of the podQOSEnsurancePolicy, and returns the corresponding podQOSEnsurancePolicy object, and an error if there is any.
func (c *FakePodQOSEnsurancePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodQOSEnsurancePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podqosensurancepoliciesResource, c.ns, name), &v1alpha1.PodQOSEnsurancePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOSEnsurancePolicy), err
}

// List takes label and field selectors, and returns the list of PodQOSEnsurancePolicies that match those selectors.
func (c *FakePodQOSEnsurancePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodQOSEnsurancePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podqosensurancepoliciesResource, podqosensurancepoliciesKind, c.ns, opts), &v1alpha1.PodQOSEnsurancePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodQOSEnsurancePolicyList{ListMeta: obj.(*v1alpha1.PodQOSEnsurancePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodQOSEnsurancePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podQOSEnsurancePolicies.
func (c *FakePodQOSEnsurancePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podqosensurancepoliciesResource, c.ns, opts))

}

// Create takes the representation of a podQOSEnsurancePolicy and creates it.  Returns the server's representation of the podQOSEnsurancePolicy, and an error, if there is any.
func (c *FakePodQOSEnsurancePolicies) Create(ctx context.Context, podQOSEnsurancePolicy *v1alpha1.PodQOSEnsurancePolicy, opts v1.CreateOptions) (result *v1alpha1.PodQOSEnsurancePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podqosensurancepoliciesResource, c.ns, podQOSEnsurancePolicy), &v1alpha1.PodQOSEnsurancePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOSEnsurancePolicy), err
}

// Update takes the representation of a podQOSEnsurancePolicy and updates it. Returns the server's representation of the podQOSEnsurancePolicy, and an error, if there is any.
func (c *FakePodQOSEnsurancePolicies) Update(ctx context.Context, podQOSEnsurancePolicy *v1alpha1.PodQOSEnsurancePolicy, opts v1.UpdateOptions) (result *v1alpha1.PodQOSEnsurancePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podqosensurancepoliciesResource, c.ns, podQOSEnsurancePolicy), &v1alpha1.PodQOSEnsurancePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOSEnsurancePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodQOSEnsurancePolicies) UpdateStatus(ctx context.Context, podQOSEnsurancePolicy *v1alpha1.PodQOSEnsurancePolicy, opts v1.UpdateOptions) (*v1alpha1.PodQOSEnsurancePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podqosensurancepoliciesResource, "status", c.ns, podQOSEnsurancePolicy), &v1alpha1.PodQOSEnsurancePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOSEnsurancePolicy), err
}

// Delete takes name of the podQOSEnsurancePolicy and deletes it. Returns an error if one occurs.
func (c *FakePodQOSEnsurancePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podqosensurancepoliciesResource, c.ns, name), &v1alpha1.PodQOSEnsurancePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodQOSEnsurancePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podqosensurancepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodQOSEnsurancePolicyList{})
	return err
}

// Patch applies the patch and returns the patched podQOSEnsurancePolicy.
func (c *FakePodQOSEnsurancePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodQOSEnsurancePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podqosensurancepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PodQOSEnsurancePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodQOSEnsurancePolicy), err
}
