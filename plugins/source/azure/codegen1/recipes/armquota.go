// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota"

func Armquota() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armquota.NewOperationClient,
		},
    
		{
			NewFunc: armquota.NewRequestStatusClient,
		},
    
		{
			NewFunc: armquota.NewUsagesClient,
		},
    
		{
			NewFunc: armquota.NewClient,
		},
    
	}
	return resources
}