// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbiembedded/armpowerbiembedded"

func Armpowerbiembedded() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpowerbiembedded.NewWorkspaceCollectionsClient,
		},
		{
			NewFunc: armpowerbiembedded.NewWorkspacesClient,
		},
		{
			NewFunc: armpowerbiembedded.NewManagementClient,
		},
	}
	return resources
}