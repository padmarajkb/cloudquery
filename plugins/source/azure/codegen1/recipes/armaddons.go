// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/addons/armaddons"

func Armaddons() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armaddons.NewOperationsClient,
		},
    
		{
			NewFunc: armaddons.NewSupportPlanTypesClient,
		},
    
	}
	return resources
}