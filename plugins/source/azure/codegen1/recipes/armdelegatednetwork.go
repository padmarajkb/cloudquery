// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"

func Armdelegatednetwork() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armdelegatednetwork.NewDelegatedSubnetServiceClient,
		},
    
		{
			NewFunc: armdelegatednetwork.NewOperationsClient,
		},
    
		{
			NewFunc: armdelegatednetwork.NewOrchestratorInstanceServiceClient,
		},
    
		{
			NewFunc: armdelegatednetwork.NewClient,
		},
    
		{
			NewFunc: armdelegatednetwork.NewControllerClient,
		},
    
	}
	return resources
}