// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ServiceCatalogControllerManagerLister helps list ServiceCatalogControllerManagers.
// All objects returned here must be treated as read-only.
type ServiceCatalogControllerManagerLister interface {
	// List lists all ServiceCatalogControllerManagers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ServiceCatalogControllerManager, err error)
	// Get retrieves the ServiceCatalogControllerManager from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ServiceCatalogControllerManager, error)
	ServiceCatalogControllerManagerListerExpansion
}

// serviceCatalogControllerManagerLister implements the ServiceCatalogControllerManagerLister interface.
type serviceCatalogControllerManagerLister struct {
	listers.ResourceIndexer[*v1.ServiceCatalogControllerManager]
}

// NewServiceCatalogControllerManagerLister returns a new ServiceCatalogControllerManagerLister.
func NewServiceCatalogControllerManagerLister(indexer cache.Indexer) ServiceCatalogControllerManagerLister {
	return &serviceCatalogControllerManagerLister{listers.New[*v1.ServiceCatalogControllerManager](indexer, v1.Resource("servicecatalogcontrollermanager"))}
}
