// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/gocrane-io/api/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// NodeQOSEnsurancePolicies returns a NodeQOSEnsurancePolicyInformer.
	NodeQOSEnsurancePolicies() NodeQOSEnsurancePolicyInformer
	// PodQOSEnsurancePolicies returns a PodQOSEnsurancePolicyInformer.
	PodQOSEnsurancePolicies() PodQOSEnsurancePolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// NodeQOSEnsurancePolicies returns a NodeQOSEnsurancePolicyInformer.
func (v *version) NodeQOSEnsurancePolicies() NodeQOSEnsurancePolicyInformer {
	return &nodeQOSEnsurancePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodQOSEnsurancePolicies returns a PodQOSEnsurancePolicyInformer.
func (v *version) PodQOSEnsurancePolicies() PodQOSEnsurancePolicyInformer {
	return &podQOSEnsurancePolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
