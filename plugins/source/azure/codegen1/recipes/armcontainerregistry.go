// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"

func Armcontainerregistry() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcontainerregistry.NewRunsClient,
		},
		{
			NewFunc: armcontainerregistry.NewTaskRunsClient,
		},
		{
			NewFunc: armcontainerregistry.NewTasksClient,
		},
		{
			NewFunc: armcontainerregistry.NewExportPipelinesClient,
		},
		{
			NewFunc: armcontainerregistry.NewReplicationsClient,
		},
		{
			NewFunc: armcontainerregistry.NewWebhooksClient,
		},
		{
			NewFunc: armcontainerregistry.NewAgentPoolsClient,
		},
		{
			NewFunc: armcontainerregistry.NewPipelineRunsClient,
		},
		{
			NewFunc: armcontainerregistry.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armcontainerregistry.NewScopeMapsClient,
		},
		{
			NewFunc: armcontainerregistry.NewTokensClient,
		},
		{
			NewFunc: armcontainerregistry.NewConnectedRegistriesClient,
		},
		{
			NewFunc: armcontainerregistry.NewImportPipelinesClient,
		},
		{
			NewFunc: armcontainerregistry.NewOperationsClient,
		},
		{
			NewFunc: armcontainerregistry.NewRegistriesClient,
		},
	}
	return resources
}