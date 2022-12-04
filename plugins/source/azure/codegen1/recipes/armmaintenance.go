// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance"

func Armmaintenance() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armmaintenance.NewConfigurationsForResourceGroupClient,
		},
		{
			NewFunc: armmaintenance.NewUpdatesClient,
		},
		{
			NewFunc: armmaintenance.NewApplyUpdateForResourceGroupClient,
		},
		{
			NewFunc: armmaintenance.NewApplyUpdatesClient,
		},
		{
			NewFunc: armmaintenance.NewOperationsClient,
		},
		{
			NewFunc: armmaintenance.NewPublicMaintenanceConfigurationsClient,
		},
		{
			NewFunc: armmaintenance.NewConfigurationsClient,
		},
		{
			NewFunc: armmaintenance.NewConfigurationAssignmentsClient,
		},
	}
	return resources
}