// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic"

func Armelastic() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armelastic.NewAllTrafficFiltersClient,
		},
		{
			NewFunc: armelastic.NewCreateAndAssociatePLFilterClient,
		},
		{
			NewFunc: armelastic.NewDeploymentInfoClient,
		},
		{
			NewFunc: armelastic.NewMonitoredResourcesClient,
		},
		{
			NewFunc: armelastic.NewVMHostClient,
		},
		{
			NewFunc: armelastic.NewDetachAndDeleteTrafficFilterClient,
		},
		{
			NewFunc: armelastic.NewExternalUserClient,
		},
		{
			NewFunc: armelastic.NewUpgradableVersionsClient,
		},
		{
			NewFunc: armelastic.NewTagRulesClient,
		},
		{
			NewFunc: armelastic.NewAssociateTrafficFilterClient,
		},
		{
			NewFunc: armelastic.NewCreateAndAssociateIPFilterClient,
		},
		{
			NewFunc: armelastic.NewDetachTrafficFilterClient,
		},
		{
			NewFunc: armelastic.NewMonitorsClient,
		},
		{
			NewFunc: armelastic.NewOperationsClient,
		},
		{
			NewFunc: armelastic.NewListAssociatedTrafficFiltersClient,
		},
		{
			NewFunc: armelastic.NewMonitorClient,
		},
		{
			NewFunc: armelastic.NewTrafficFiltersClient,
		},
		{
			NewFunc: armelastic.NewVMCollectionClient,
		},
		{
			NewFunc: armelastic.NewVMIngestionClient,
		},
	}
	return resources
}