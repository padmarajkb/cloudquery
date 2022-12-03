// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"

func Armappcomplianceautomation() []*Resource {
	resources := []*Resource{
    
		{
			NewFunc: armappcomplianceautomation.NewReportClient,
		},
    
		{
			NewFunc: armappcomplianceautomation.NewReportsClient,
		},
    
		{
			NewFunc: armappcomplianceautomation.NewSnapshotClient,
		},
    
		{
			NewFunc: armappcomplianceautomation.NewSnapshotsClient,
		},
    
		{
			NewFunc: armappcomplianceautomation.NewOperationsClient,
		},
    
	}
	return resources
}