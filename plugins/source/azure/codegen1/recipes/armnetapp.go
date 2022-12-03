// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp"

func Armnetapp() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armnetapp.NewVolumeGroupsClient,
		},
    
		{
			NewFunc: armnetapp.NewVolumeQuotaRulesClient,
		},
    
		{
			NewFunc: armnetapp.NewVolumesClient,
		},
    
		{
			NewFunc: armnetapp.NewAccountsClient,
		},
    
		{
			NewFunc: armnetapp.NewOperationsClient,
		},
    
		{
			NewFunc: armnetapp.NewPoolsClient,
		},
    
		{
			NewFunc: armnetapp.NewVaultsClient,
		},
    
		{
			NewFunc: armnetapp.NewAccountBackupsClient,
		},
    
		{
			NewFunc: armnetapp.NewBackupsClient,
		},
    
		{
			NewFunc: armnetapp.NewResourceClient,
		},
    
		{
			NewFunc: armnetapp.NewBackupPoliciesClient,
		},
    
		{
			NewFunc: armnetapp.NewSnapshotPoliciesClient,
		},
    
		{
			NewFunc: armnetapp.NewSnapshotsClient,
		},
    
		{
			NewFunc: armnetapp.NewResourceQuotaLimitsClient,
		},
    
		{
			NewFunc: armnetapp.NewSubvolumesClient,
		},
    
	}
	return resources
}