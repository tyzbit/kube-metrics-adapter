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

package v1

import (
	v1 "github.com/zalando-incubator/kube-metrics-adapter/pkg/apis/zalando.org/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScalingScheduleLister helps list ScalingSchedules.
// All objects returned here must be treated as read-only.
type ScalingScheduleLister interface {
	// List lists all ScalingSchedules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ScalingSchedule, err error)
	// ScalingSchedules returns an object that can list and get ScalingSchedules.
	ScalingSchedules(namespace string) ScalingScheduleNamespaceLister
	ScalingScheduleListerExpansion
}

// scalingScheduleLister implements the ScalingScheduleLister interface.
type scalingScheduleLister struct {
	indexer cache.Indexer
}

// NewScalingScheduleLister returns a new ScalingScheduleLister.
func NewScalingScheduleLister(indexer cache.Indexer) ScalingScheduleLister {
	return &scalingScheduleLister{indexer: indexer}
}

// List lists all ScalingSchedules in the indexer.
func (s *scalingScheduleLister) List(selector labels.Selector) (ret []*v1.ScalingSchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScalingSchedule))
	})
	return ret, err
}

// ScalingSchedules returns an object that can list and get ScalingSchedules.
func (s *scalingScheduleLister) ScalingSchedules(namespace string) ScalingScheduleNamespaceLister {
	return scalingScheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScalingScheduleNamespaceLister helps list and get ScalingSchedules.
// All objects returned here must be treated as read-only.
type ScalingScheduleNamespaceLister interface {
	// List lists all ScalingSchedules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ScalingSchedule, err error)
	// Get retrieves the ScalingSchedule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ScalingSchedule, error)
	ScalingScheduleNamespaceListerExpansion
}

// scalingScheduleNamespaceLister implements the ScalingScheduleNamespaceLister
// interface.
type scalingScheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScalingSchedules in the indexer for a given namespace.
func (s scalingScheduleNamespaceLister) List(selector labels.Selector) (ret []*v1.ScalingSchedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ScalingSchedule))
	})
	return ret, err
}

// Get retrieves the ScalingSchedule from the indexer for a given namespace and name.
func (s scalingScheduleNamespaceLister) Get(name string) (*v1.ScalingSchedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("scalingschedule"), name)
	}
	return obj.(*v1.ScalingSchedule), nil
}
