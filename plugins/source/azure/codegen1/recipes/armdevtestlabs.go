// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"

func Armdevtestlabs() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdevtestlabs.NewServiceRunnersClient,
		},
		{
			NewFunc: armdevtestlabs.NewSchedulesClient,
		},
		{
			NewFunc: armdevtestlabs.NewServiceFabricsClient,
		},
		{
			NewFunc: armdevtestlabs.NewSecretsClient,
		},
		{
			NewFunc: armdevtestlabs.NewProviderOperationsClient,
		},
		{
			NewFunc: armdevtestlabs.NewArtifactSourcesClient,
		},
		{
			NewFunc: armdevtestlabs.NewGalleryImagesClient,
		},
		{
			NewFunc: armdevtestlabs.NewPoliciesClient,
		},
		{
			NewFunc: armdevtestlabs.NewPolicySetsClient,
		},
		{
			NewFunc: armdevtestlabs.NewArmTemplatesClient,
		},
		{
			NewFunc: armdevtestlabs.NewArtifactsClient,
		},
		{
			NewFunc: armdevtestlabs.NewEnvironmentsClient,
		},
		{
			NewFunc: armdevtestlabs.NewGlobalSchedulesClient,
		},
		{
			NewFunc: armdevtestlabs.NewUsersClient,
		},
		{
			NewFunc: armdevtestlabs.NewDisksClient,
		},
		{
			NewFunc: armdevtestlabs.NewNotificationChannelsClient,
		},
		{
			NewFunc: armdevtestlabs.NewOperationsClient,
		},
		{
			NewFunc: armdevtestlabs.NewVirtualMachineSchedulesClient,
		},
		{
			NewFunc: armdevtestlabs.NewLabsClient,
		},
		{
			NewFunc: armdevtestlabs.NewCostsClient,
		},
		{
			NewFunc: armdevtestlabs.NewVirtualNetworksClient,
		},
		{
			NewFunc: armdevtestlabs.NewVirtualMachinesClient,
		},
		{
			NewFunc: armdevtestlabs.NewServiceFabricSchedulesClient,
		},
		{
			NewFunc: armdevtestlabs.NewCustomImagesClient,
		},
		{
			NewFunc: armdevtestlabs.NewFormulasClient,
		},
	}
	return resources
}