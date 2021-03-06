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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/bradmwilliams/release-payload-controller/pkg/apis/release/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReleasePayloadLister helps list ReleasePayloads.
// All objects returned here must be treated as read-only.
type ReleasePayloadLister interface {
	// List lists all ReleasePayloads in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReleasePayload, err error)
	// ReleasePayloads returns an object that can list and get ReleasePayloads.
	ReleasePayloads(namespace string) ReleasePayloadNamespaceLister
	ReleasePayloadListerExpansion
}

// releasePayloadLister implements the ReleasePayloadLister interface.
type releasePayloadLister struct {
	indexer cache.Indexer
}

// NewReleasePayloadLister returns a new ReleasePayloadLister.
func NewReleasePayloadLister(indexer cache.Indexer) ReleasePayloadLister {
	return &releasePayloadLister{indexer: indexer}
}

// List lists all ReleasePayloads in the indexer.
func (s *releasePayloadLister) List(selector labels.Selector) (ret []*v1alpha1.ReleasePayload, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReleasePayload))
	})
	return ret, err
}

// ReleasePayloads returns an object that can list and get ReleasePayloads.
func (s *releasePayloadLister) ReleasePayloads(namespace string) ReleasePayloadNamespaceLister {
	return releasePayloadNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReleasePayloadNamespaceLister helps list and get ReleasePayloads.
// All objects returned here must be treated as read-only.
type ReleasePayloadNamespaceLister interface {
	// List lists all ReleasePayloads in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReleasePayload, err error)
	// Get retrieves the ReleasePayload from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ReleasePayload, error)
	ReleasePayloadNamespaceListerExpansion
}

// releasePayloadNamespaceLister implements the ReleasePayloadNamespaceLister
// interface.
type releasePayloadNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReleasePayloads in the indexer for a given namespace.
func (s releasePayloadNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReleasePayload, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReleasePayload))
	})
	return ret, err
}

// Get retrieves the ReleasePayload from the indexer for a given namespace and name.
func (s releasePayloadNamespaceLister) Get(name string) (*v1alpha1.ReleasePayload, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("releasepayload"), name)
	}
	return obj.(*v1alpha1.ReleasePayload), nil
}
