// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"

func Armsupport() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armsupport.NewCommunicationsClient,
		},
    
		{
			NewFunc: armsupport.NewOperationsClient,
		},
    
		{
			NewFunc: armsupport.NewProblemClassificationsClient,
		},
    
		{
			NewFunc: armsupport.NewServicesClient,
		},
    
		{
			NewFunc: armsupport.NewTicketsClient,
		},
    
	}
	return resources
}