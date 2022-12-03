// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationsmanagement/armoperationsmanagement"

func Armoperationsmanagement() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armoperationsmanagement.NewManagementConfigurationsClient,
		},
    
		{
			NewFunc: armoperationsmanagement.NewSolutionsClient,
		},
    
		{
			NewFunc: armoperationsmanagement.NewManagementAssociationsClient,
		},
    
		{
			NewFunc: armoperationsmanagement.NewOperationsClient,
		},
    
	}
	return resources
}