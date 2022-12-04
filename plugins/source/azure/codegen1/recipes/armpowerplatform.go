// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerplatform/armpowerplatform"

func Armpowerplatform() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpowerplatform.NewEnterprisePoliciesClient,
		},
		{
			NewFunc: armpowerplatform.NewOperationsClient,
		},
		{
			NewFunc: armpowerplatform.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armpowerplatform.NewPrivateLinkResourcesClient,
		},
		{
			NewFunc: armpowerplatform.NewAccountsClient,
		},
	}
	return resources
}