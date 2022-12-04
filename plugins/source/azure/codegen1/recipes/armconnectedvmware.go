// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/connectedvmware/armconnectedvmware"

func Armconnectedvmware() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armconnectedvmware.NewResourcePoolsClient,
		},
		{
			NewFunc: armconnectedvmware.NewVirtualNetworksClient,
		},
		{
			NewFunc: armconnectedvmware.NewHostsClient,
		},
		{
			NewFunc: armconnectedvmware.NewInventoryItemsClient,
		},
		{
			NewFunc: armconnectedvmware.NewMachineExtensionsClient,
		},
		{
			NewFunc: armconnectedvmware.NewOperationsClient,
		},
		{
			NewFunc: armconnectedvmware.NewVirtualMachinesClient,
		},
		{
			NewFunc: armconnectedvmware.NewClustersClient,
		},
		{
			NewFunc: armconnectedvmware.NewDatastoresClient,
		},
		{
			NewFunc: armconnectedvmware.NewGuestAgentsClient,
		},
		{
			NewFunc: armconnectedvmware.NewHybridIdentityMetadataClient,
		},
		{
			NewFunc: armconnectedvmware.NewVCentersClient,
		},
		{
			NewFunc: armconnectedvmware.NewVirtualMachineTemplatesClient,
		},
	}
	return resources
}