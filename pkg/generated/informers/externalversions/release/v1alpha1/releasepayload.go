/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	releasev1alpha1 "github.com/bradmwilliams/release-payload-controller/pkg/apis/release/v1alpha1"
	versioned "github.com/bradmwilliams/release-payload-controller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/bradmwilliams/release-payload-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bradmwilliams/release-payload-controller/pkg/generated/listers/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReleasePayloadInformer provides access to a shared informer and lister for
// ReleasePayloads.
type ReleasePayloadInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReleasePayloadLister
}

type releasePayloadInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReleasePayloadInformer constructs a new informer for ReleasePayload type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReleasePayloadInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReleasePayloadInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReleasePayloadInformer constructs a new informer for ReleasePayload type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReleasePayloadInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ReleasePayloads(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ReleasePayloads(namespace).Watch(context.TODO(), options)
			},
		},
		&releasev1alpha1.ReleasePayload{},
		resyncPeriod,
		indexers,
	)
}

func (f *releasePayloadInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReleasePayloadInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *releasePayloadInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&releasev1alpha1.ReleasePayload{}, f.defaultInformer)
}

func (f *releasePayloadInformer) Lister() v1alpha1.ReleasePayloadLister {
	return v1alpha1.NewReleasePayloadLister(f.Informer().GetIndexer())
}