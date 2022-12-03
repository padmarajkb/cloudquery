// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"

func Armmachinelearningservices() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armmachinelearningservices.NewComputeClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewUsagesClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewWorkspaceFeaturesClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewQuotasClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewWorkspacesClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewVirtualMachineSizesClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewWorkspaceConnectionsClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewWorkspaceSKUsClient,
		},
    
		{
			NewFunc: armmachinelearningservices.NewOperationsClient,
		},
    
	}
	return resources
}