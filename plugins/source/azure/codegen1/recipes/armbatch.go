// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"

func Armbatch() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armbatch.NewOperationsClient,
		},
		{
			NewFunc: armbatch.NewPoolClient,
		},
		{
			NewFunc: armbatch.NewApplicationClient,
		},
		{
			NewFunc: armbatch.NewApplicationPackageClient,
		},
		{
			NewFunc: armbatch.NewAccountClient,
		},
		{
			NewFunc: armbatch.NewLocationClient,
		},
		{
			NewFunc: armbatch.NewPrivateLinkResourceClient,
		},
		{
			NewFunc: armbatch.NewCertificateClient,
		},
		{
			NewFunc: armbatch.NewPrivateEndpointConnectionClient,
		},
	}
	return resources
}