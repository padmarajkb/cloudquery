// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog"

func Armdatadog() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armdatadog.NewOperationsClient,
		},
		{
			NewFunc: armdatadog.NewSingleSignOnConfigurationsClient,
		},
		{
			NewFunc: armdatadog.NewMarketplaceAgreementsClient,
		},
		{
			NewFunc: armdatadog.NewMonitorsClient,
		},
		{
			NewFunc: armdatadog.NewTagRulesClient,
		},
	}
	return resources
}