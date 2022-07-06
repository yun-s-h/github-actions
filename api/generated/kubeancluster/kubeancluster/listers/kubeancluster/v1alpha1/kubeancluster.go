// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubean.io/api/apis/kubeancluster/v1alpha1"
)

// KuBeanClusterLister helps list KuBeanClusters.
// All objects returned here must be treated as read-only.
type KuBeanClusterLister interface {
	// List lists all KuBeanClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KuBeanCluster, err error)
	// Get retrieves the KuBeanCluster from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KuBeanCluster, error)
	KuBeanClusterListerExpansion
}

// kuBeanClusterLister implements the KuBeanClusterLister interface.
type kuBeanClusterLister struct {
	indexer cache.Indexer
}

// NewKuBeanClusterLister returns a new KuBeanClusterLister.
func NewKuBeanClusterLister(indexer cache.Indexer) KuBeanClusterLister {
	return &kuBeanClusterLister{indexer: indexer}
}

// List lists all KuBeanClusters in the indexer.
func (s *kuBeanClusterLister) List(selector labels.Selector) (ret []*v1alpha1.KuBeanCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KuBeanCluster))
	})
	return ret, err
}

// Get retrieves the KuBeanCluster from the index for a given name.
func (s *kuBeanClusterLister) Get(name string) (*v1alpha1.KuBeanCluster, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubeancluster"), name)
	}
	return obj.(*v1alpha1.KuBeanCluster), nil
}
