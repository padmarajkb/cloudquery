// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hanaonazure/armhanaonazure"

func Armhanaonazure() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armhanaonazure.NewSapMonitorsClient,
		},
		{
			NewFunc: armhanaonazure.NewOperationsClient,
		},
		{
			NewFunc: armhanaonazure.NewProviderInstancesClient,
		},
	}
	return resources
}