// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dns/armdns"

func Armdns() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armdns.NewZonesClient,
		},
    
		{
			NewFunc: armdns.NewRecordSetsClient,
		},
    
		{
			NewFunc: armdns.NewResourceReferenceClient,
		},
    
	}
	return resources
}