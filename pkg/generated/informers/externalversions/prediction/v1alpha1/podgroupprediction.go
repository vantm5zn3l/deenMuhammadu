// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/gocrane/api/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/gocrane/api/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/gocrane/api/pkg/generated/listers/prediction/v1alpha1"
	predictionv1alpha1 "github.com/gocrane/api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodGroupPredictionInformer provides access to a shared informer and lister for
// PodGroupPredictions.
type PodGroupPredictionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodGroupPredictionLister
}

type podGroupPredictionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodGroupPredictionInformer constructs a new informer for PodGroupPrediction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodGroupPredictionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodGroupPredictionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodGroupPredictionInformer constructs a new informer for PodGroupPrediction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodGroupPredictionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PredictionV1alpha1().PodGroupPredictions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PredictionV1alpha1().PodGroupPredictions(namespace).Watch(context.TODO(), options)
			},
		},
		&predictionv1alpha1.PodGroupPrediction{},
		resyncPeriod,
		indexers,
	)
}

func (f *podGroupPredictionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodGroupPredictionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podGroupPredictionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&predictionv1alpha1.PodGroupPrediction{}, f.defaultInformer)
}

func (f *podGroupPredictionInformer) Lister() v1alpha1.PodGroupPredictionLister {
	return v1alpha1.NewPodGroupPredictionLister(f.Informer().GetIndexer())
}
