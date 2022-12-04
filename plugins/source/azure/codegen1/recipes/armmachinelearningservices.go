// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"

func Armmachinelearningservices() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmachinelearningservices.NewWorkspaceFeaturesClient,
		},
		{
			NewFunc: armmachinelearningservices.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armmachinelearningservices.NewPrivateLinkResourcesClient,
		},
		{
			NewFunc: armmachinelearningservices.NewQuotasClient,
		},
		{
			NewFunc: armmachinelearningservices.NewWorkspaceSKUsClient,
		},
		{
			NewFunc: armmachinelearningservices.NewComputeClient,
		},
		{
			NewFunc: armmachinelearningservices.NewUsagesClient,
		},
		{
			NewFunc: armmachinelearningservices.NewVirtualMachineSizesClient,
		},
		{
			NewFunc: armmachinelearningservices.NewWorkspaceConnectionsClient,
		},
		{
			NewFunc: armmachinelearningservices.NewOperationsClient,
		},
		{
			NewFunc: armmachinelearningservices.NewWorkspacesClient,
		},
	}
	return resources
}