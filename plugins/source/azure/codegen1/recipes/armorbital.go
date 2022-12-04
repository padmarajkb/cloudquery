// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital"

func Armorbital() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armorbital.NewAvailableGroundStationsClient,
		},
		{
			NewFunc: armorbital.NewContactsClient,
		},
		{
			NewFunc: armorbital.NewSpacecraftsClient,
		},
		{
			NewFunc: armorbital.NewOperationsClient,
		},
		{
			NewFunc: armorbital.NewOperationsResultsClient,
		},
		{
			NewFunc: armorbital.NewContactProfilesClient,
		},
	}
	return resources
}