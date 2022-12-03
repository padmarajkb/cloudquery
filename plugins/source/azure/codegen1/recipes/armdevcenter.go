// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter"

func Armdevcenter() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armdevcenter.NewGalleriesClient,
		},
    
		{
			NewFunc: armdevcenter.NewNetworkConnectionsClient,
		},
    
		{
			NewFunc: armdevcenter.NewOperationsClient,
		},
    
		{
			NewFunc: armdevcenter.NewProjectsClient,
		},
    
		{
			NewFunc: armdevcenter.NewDevBoxDefinitionsClient,
		},
    
		{
			NewFunc: armdevcenter.NewDevCentersClient,
		},
    
		{
			NewFunc: armdevcenter.NewEnvironmentTypesClient,
		},
    
		{
			NewFunc: armdevcenter.NewProjectAllowedEnvironmentTypesClient,
		},
    
		{
			NewFunc: armdevcenter.NewAttachedNetworksClient,
		},
    
		{
			NewFunc: armdevcenter.NewUsagesClient,
		},
    
		{
			NewFunc: armdevcenter.NewCatalogsClient,
		},
    
		{
			NewFunc: armdevcenter.NewOperationStatusesClient,
		},
    
		{
			NewFunc: armdevcenter.NewPoolsClient,
		},
    
		{
			NewFunc: armdevcenter.NewProjectEnvironmentTypesClient,
		},
    
		{
			NewFunc: armdevcenter.NewSchedulesClient,
		},
    
		{
			NewFunc: armdevcenter.NewSKUsClient,
		},
    
		{
			NewFunc: armdevcenter.NewImagesClient,
		},
    
		{
			NewFunc: armdevcenter.NewImageVersionsClient,
		},
    
	}
	return resources
}