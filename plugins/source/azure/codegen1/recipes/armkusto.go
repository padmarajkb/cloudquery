// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"

func Armkusto() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armkusto.NewDataConnectionsClient,
		},
		{
			NewFunc: armkusto.NewOperationsResultsClient,
		},
		{
			NewFunc: armkusto.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armkusto.NewClusterPrincipalAssignmentsClient,
		},
		{
			NewFunc: armkusto.NewClustersClient,
		},
		{
			NewFunc: armkusto.NewDatabasePrincipalAssignmentsClient,
		},
		{
			NewFunc: armkusto.NewManagedPrivateEndpointsClient,
		},
		{
			NewFunc: armkusto.NewOperationsResultsLocationClient,
		},
		{
			NewFunc: armkusto.NewAttachedDatabaseConfigurationsClient,
		},
		{
			NewFunc: armkusto.NewDatabasesClient,
		},
		{
			NewFunc: armkusto.NewOperationsClient,
		},
		{
			NewFunc: armkusto.NewPrivateLinkResourcesClient,
		},
		{
			NewFunc: armkusto.NewScriptsClient,
		},
	}
	return resources
}