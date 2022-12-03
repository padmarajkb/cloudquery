// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportClient,
		},
    
		{
			NewFunc: armcostmanagement.NewQueryClient,
		},
    
		{
			NewFunc: armcostmanagement.NewAlertsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewForecastClient,
		},
    
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationResultsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationStatusClient,
		},
    
		{
			NewFunc: armcostmanagement.NewOperationsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewViewsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewDimensionsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewExportsClient,
		},
    
		{
			NewFunc: armcostmanagement.NewGenerateReservationDetailsReportClient,
		},
    
	}
	return resources
}