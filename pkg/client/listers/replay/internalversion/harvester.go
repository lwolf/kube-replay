/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	replay "github.com/lwolf/kube-replay/pkg/apis/replay"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HarvesterLister helps list Harvesters.
type HarvesterLister interface {
	// List lists all Harvesters in the indexer.
	List(selector labels.Selector) (ret []*replay.Harvester, err error)
	// Harvesters returns an object that can list and get Harvesters.
	Harvesters(namespace string) HarvesterNamespaceLister
	HarvesterListerExpansion
}

// harvesterLister implements the HarvesterLister interface.
type harvesterLister struct {
	indexer cache.Indexer
}

// NewHarvesterLister returns a new HarvesterLister.
func NewHarvesterLister(indexer cache.Indexer) HarvesterLister {
	return &harvesterLister{indexer: indexer}
}

// List lists all Harvesters in the indexer.
func (s *harvesterLister) List(selector labels.Selector) (ret []*replay.Harvester, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*replay.Harvester))
	})
	return ret, err
}

// Harvesters returns an object that can list and get Harvesters.
func (s *harvesterLister) Harvesters(namespace string) HarvesterNamespaceLister {
	return harvesterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HarvesterNamespaceLister helps list and get Harvesters.
type HarvesterNamespaceLister interface {
	// List lists all Harvesters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*replay.Harvester, err error)
	// Get retrieves the Harvester from the indexer for a given namespace and name.
	Get(name string) (*replay.Harvester, error)
	HarvesterNamespaceListerExpansion
}

// harvesterNamespaceLister implements the HarvesterNamespaceLister
// interface.
type harvesterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Harvesters in the indexer for a given namespace.
func (s harvesterNamespaceLister) List(selector labels.Selector) (ret []*replay.Harvester, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*replay.Harvester))
	})
	return ret, err
}

// Get retrieves the Harvester from the indexer for a given namespace and name.
func (s harvesterNamespaceLister) Get(name string) (*replay.Harvester, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(replay.Resource("harvester"), name)
	}
	return obj.(*replay.Harvester), nil
}
