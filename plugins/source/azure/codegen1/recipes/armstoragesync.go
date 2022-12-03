// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"

func Armstoragesync() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armstoragesync.NewCloudEndpointsClient,
		},
    
		{
			NewFunc: armstoragesync.NewServicesClient,
		},
    
		{
			NewFunc: armstoragesync.NewOperationsClient,
		},
    
		{
			NewFunc: armstoragesync.NewPrivateEndpointConnectionsClient,
		},
    
		{
			NewFunc: armstoragesync.NewRegisteredServersClient,
		},
    
		{
			NewFunc: armstoragesync.NewServerEndpointsClient,
		},
    
		{
			NewFunc: armstoragesync.NewMicrosoftStorageSyncClient,
		},
    
		{
			NewFunc: armstoragesync.NewWorkflowsClient,
		},
    
		{
			NewFunc: armstoragesync.NewOperationStatusClient,
		},
    
		{
			NewFunc: armstoragesync.NewPrivateLinkResourcesClient,
		},
    
		{
			NewFunc: armstoragesync.NewSyncGroupsClient,
		},
    
	}
	return resources
}