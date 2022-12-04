// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu"

func Armwindowsesu() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armwindowsesu.NewMultipleActivationKeysClient,
		},
		{
			NewFunc: armwindowsesu.NewOperationsClient,
		},
	}
	return resources
}