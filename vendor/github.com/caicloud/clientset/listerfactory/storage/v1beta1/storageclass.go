/*
Copyright 2019 caicloud authors. All rights reserved.
*/

// Code generated by listerfactory-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "github.com/caicloud/clientset/listerfactory/internalinterfaces"
	storagev1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	kubernetes "k8s.io/client-go/kubernetes"
	v1beta1 "k8s.io/client-go/listers/storage/v1beta1"
)

var _ v1beta1.StorageClassLister = &storageClassLister{}

// storageClassLister implements the StorageClassLister interface.
type storageClassLister struct {
	client           kubernetes.Interface
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStorageClassLister returns a new StorageClassLister.
func NewStorageClassLister(client kubernetes.Interface) v1beta1.StorageClassLister {
	return NewFilteredStorageClassLister(client, nil)
}

func NewFilteredStorageClassLister(client kubernetes.Interface, tweakListOptions internalinterfaces.TweakListOptionsFunc) v1beta1.StorageClassLister {
	return &storageClassLister{
		client:           client,
		tweakListOptions: tweakListOptions,
	}
}

// List lists all StorageClasses in the indexer.
func (s *storageClassLister) List(selector labels.Selector) (ret []*storagev1beta1.StorageClass, err error) {
	listopt := v1.ListOptions{
		LabelSelector: selector.String(),
	}
	if s.tweakListOptions != nil {
		s.tweakListOptions(&listopt)
	}
	list, err := s.client.StorageV1beta1().StorageClasses().List(listopt)
	if err != nil {
		return nil, err
	}
	for i := range list.Items {
		ret = append(ret, &list.Items[i])
	}
	return ret, nil
}

// Get retrieves the StorageClass from the index for a given name.
func (s *storageClassLister) Get(name string) (*storagev1beta1.StorageClass, error) {
	return s.client.StorageV1beta1().StorageClasses().Get(name, v1.GetOptions{})
}
