// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"

func Armcostmanagement() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armcostmanagement.NewExportsClient,
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationStatusClient,
		},
		{
			NewFunc: armcostmanagement.NewGenerateReservationDetailsReportClient,
		},
		{
			NewFunc: armcostmanagement.NewOperationsClient,
		},
		{
			NewFunc: armcostmanagement.NewAlertsClient,
		},
		{
			NewFunc: armcostmanagement.NewViewsClient,
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportClient,
		},
		{
			NewFunc: armcostmanagement.NewGenerateDetailedCostReportOperationResultsClient,
		},
		{
			NewFunc: armcostmanagement.NewQueryClient,
		},
		{
			NewFunc: armcostmanagement.NewForecastClient,
		},
		{
			NewFunc: armcostmanagement.NewDimensionsClient,
		},
	}
	return resources
}