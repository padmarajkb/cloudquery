// Code generated by codegen; DO NOT EDIT.
package recipes

import "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"

func Armpostgresql() []*Resource {
	resources := []*Resource{
		{
			NewFunc: armpostgresql.NewPrivateEndpointConnectionsClient,
		},
		{
			NewFunc: armpostgresql.NewPrivateLinkResourcesClient,
		},
		{
			NewFunc: armpostgresql.NewServerAdministratorsClient,
		},
		{
			NewFunc: armpostgresql.NewServerBasedPerformanceTierClient,
		},
		{
			NewFunc: armpostgresql.NewServerParametersClient,
		},
		{
			NewFunc: armpostgresql.NewServerSecurityAlertPoliciesClient,
		},
		{
			NewFunc: armpostgresql.NewVirtualNetworkRulesClient,
		},
		{
			NewFunc: armpostgresql.NewLocationBasedPerformanceTierClient,
		},
		{
			NewFunc: armpostgresql.NewServerKeysClient,
		},
		{
			NewFunc: armpostgresql.NewServersClient,
		},
		{
			NewFunc: armpostgresql.NewDatabasesClient,
		},
		{
			NewFunc: armpostgresql.NewFirewallRulesClient,
		},
		{
			NewFunc: armpostgresql.NewOperationsClient,
		},
		{
			NewFunc: armpostgresql.NewCheckNameAvailabilityClient,
		},
		{
			NewFunc: armpostgresql.NewConfigurationsClient,
		},
		{
			NewFunc: armpostgresql.NewLogFilesClient,
		},
		{
			NewFunc: armpostgresql.NewRecoverableServersClient,
		},
		{
			NewFunc: armpostgresql.NewReplicasClient,
		},
	}
	return resources
}