// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/gocrane-io/api/ensurance/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeQOSEnsurancePolicyLister helps list NodeQOSEnsurancePolicies.
// All objects returned here must be treated as read-only.
type NodeQOSEnsurancePolicyLister interface {
	// List lists all NodeQOSEnsurancePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeQOSEnsurancePolicy, err error)
	// NodeQOSEnsurancePolicies returns an object that can list and get NodeQOSEnsurancePolicies.
	NodeQOSEnsurancePolicies(namespace string) NodeQOSEnsurancePolicyNamespaceLister
	NodeQOSEnsurancePolicyListerExpansion
}

// nodeQOSEnsurancePolicyLister implements the NodeQOSEnsurancePolicyLister interface.
type nodeQOSEnsurancePolicyLister struct {
	indexer cache.Indexer
}

// NewNodeQOSEnsurancePolicyLister returns a new NodeQOSEnsurancePolicyLister.
func NewNodeQOSEnsurancePolicyLister(indexer cache.Indexer) NodeQOSEnsurancePolicyLister {
	return &nodeQOSEnsurancePolicyLister{indexer: indexer}
}

// List lists all NodeQOSEnsurancePolicies in the indexer.
func (s *nodeQOSEnsurancePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.NodeQOSEnsurancePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeQOSEnsurancePolicy))
	})
	return ret, err
}

// NodeQOSEnsurancePolicies returns an object that can list and get NodeQOSEnsurancePolicies.
func (s *nodeQOSEnsurancePolicyLister) NodeQOSEnsurancePolicies(namespace string) NodeQOSEnsurancePolicyNamespaceLister {
	return nodeQOSEnsurancePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeQOSEnsurancePolicyNamespaceLister helps list and get NodeQOSEnsurancePolicies.
// All objects returned here must be treated as read-only.
type NodeQOSEnsurancePolicyNamespaceLister interface {
	// List lists all NodeQOSEnsurancePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NodeQOSEnsurancePolicy, err error)
	// Get retrieves the NodeQOSEnsurancePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NodeQOSEnsurancePolicy, error)
	NodeQOSEnsurancePolicyNamespaceListerExpansion
}

// nodeQOSEnsurancePolicyNamespaceLister implements the NodeQOSEnsurancePolicyNamespaceLister
// interface.
type nodeQOSEnsurancePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeQOSEnsurancePolicies in the indexer for a given namespace.
func (s nodeQOSEnsurancePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NodeQOSEnsurancePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodeQOSEnsurancePolicy))
	})
	return ret, err
}

// Get retrieves the NodeQOSEnsurancePolicy from the indexer for a given namespace and name.
func (s nodeQOSEnsurancePolicyNamespaceLister) Get(name string) (*v1alpha1.NodeQOSEnsurancePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodeqosensurancepolicy"), name)
	}
	return obj.(*v1alpha1.NodeQOSEnsurancePolicy), nil
}
