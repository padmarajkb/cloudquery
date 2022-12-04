// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/attestation/armattestation"

func Armattestation() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armattestation.NewOperationsClient,
		},
		{
			NewFunc: armattestation.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armattestation.NewProvidersClient,
		},
	}
	return resources
}