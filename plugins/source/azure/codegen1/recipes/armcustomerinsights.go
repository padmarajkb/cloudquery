// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"

func Armcustomerinsights() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armcustomerinsights.NewAuthorizationPoliciesClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewConnectorMappingsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewConnectorsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewOperationsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewViewsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewWidgetTypesClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewInteractionsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewRelationshipsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewLinksClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewPredictionsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewProfilesClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewRelationshipLinksClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewRoleAssignmentsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewHubsClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewImagesClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewKpiClient,
		},
    
		{
			NewFunc: armcustomerinsights.NewRolesClient,
		},
    
	}
	return resources
}