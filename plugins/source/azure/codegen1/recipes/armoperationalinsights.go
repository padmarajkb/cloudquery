// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"

func Armoperationalinsights() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armoperationalinsights.NewSchemaClient,
		},
		{
			NewFunc: armoperationalinsights.NewWorkspacesClient,
		},
		{
			NewFunc: armoperationalinsights.NewDeletedWorkspacesClient,
		},
		{
			NewFunc: armoperationalinsights.NewDataSourcesClient,
		},
		{
			NewFunc: armoperationalinsights.NewGatewaysClient,
		},
		{
			NewFunc: armoperationalinsights.NewIntelligencePacksClient,
		},
		{
			NewFunc: armoperationalinsights.NewLinkedServicesClient,
		},
		{
			NewFunc: armoperationalinsights.NewManagementGroupsClient,
		},
		{
			NewFunc: armoperationalinsights.NewOperationsClient,
		},
		{
			NewFunc: armoperationalinsights.NewAvailableServiceTiersClient,
		},
		{
			NewFunc: armoperationalinsights.NewSavedSearchesClient,
		},
		{
			NewFunc: armoperationalinsights.NewWorkspacePurgeClient,
		},
		{
			NewFunc: armoperationalinsights.NewStorageInsightConfigsClient,
		},
		{
			NewFunc: armoperationalinsights.NewTablesClient,
		},
		{
			NewFunc: armoperationalinsights.NewUsagesClient,
		},
		{
			NewFunc: armoperationalinsights.NewDataExportsClient,
		},
		{
			NewFunc: armoperationalinsights.NewSharedKeysClient,
		},
		{
			NewFunc: armoperationalinsights.NewOperationStatusesClient,
		},
		{
			NewFunc: armoperationalinsights.NewClustersClient,
		},
		{
			NewFunc: armoperationalinsights.NewLinkedStorageAccountsClient,
		},
	}
	return resources
}