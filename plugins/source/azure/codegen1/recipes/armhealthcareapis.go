// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"

func Armhealthcareapis() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhealthcareapis.NewDicomServicesClient,
		},
		{
			NewFunc: armhealthcareapis.NewFhirServicesClient,
		},
		{
			NewFunc: armhealthcareapis.NewServicesClient,
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateLinkResourcesClient,
		},
		{
			NewFunc: armhealthcareapis.NewFhirDestinationsClient,
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorFhirDestinationClient,
		},
		{
			NewFunc: armhealthcareapis.NewOperationResultsClient,
		},
		{
			NewFunc: armhealthcareapis.NewOperationsClient,
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacePrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armhealthcareapis.NewWorkspacesClient,
		},
		{
			NewFunc: armhealthcareapis.NewIotConnectorsClient,
		},
		{
			NewFunc: armhealthcareapis.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armhealthcareapis.NewPrivateLinkResourcesClient,
		},
	}
	return resources
}