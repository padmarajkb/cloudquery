// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/iotcentral/armiotcentral"

func Armiotcentral() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armiotcentral.NewOperationsClient,
		},
		{
			NewFunc: armiotcentral.NewAppsClient,
		},
	}
	return resources
}