// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/edgeorder/armedgeorder"

func Armedgeorder() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armedgeorder.NewManagementClient,
		},
    
	}
	return resources
}