// Code generated by codegen; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armadvisor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armbatch"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armcdn"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armcompute"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armcontainerregistry"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armcontainerservice"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armcosmos"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armdatalakeanalytics"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armdatalakestore"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armeventhub"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armfrontdoor"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armkeyvault"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armlogic"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armmariadb"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armmysql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armnetwork"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armpostgresql"
	"github.com/cloudquery/cloudquery/plugins/source/azure/resources/services/armredis"
	"github.com/cloudquery/plugin-sdk/schema"
)

func PluginAutoGeneratedTables() []*schema.Table {
	return []*schema.Table{
		armadvisor.OperationEntities(),
		armadvisor.ConfigData(),
		armbatch.Accounts(),
		armcdn.Profiles(),
		armcompute.Disks(),
		armcompute.VirtualMachines(),
		armcompute.VirtualMachineScaleSets(),
		armcontainerregistry.Registries(),
		armcontainerservice.ManagedClusters(),
		armcosmos.Accounts(),
		armdatalakestore.Accounts(),
		armdatalakeanalytics.Accounts(),
		armeventhub.Namespaces(),
		armeventhub.NetworkRuleSets(),
		armfrontdoor.Doors(),
		armkeyvault.Vaults(),
		armlogic.Workflows(),
		armmariadb.Servers(),
		armmysql.Servers(),
		armnetwork.VirtualNetworks(),
		armnetwork.SecurityGroups(),
		armnetwork.Interfaces(),
		armnetwork.Watchers(),
		armpostgresql.Servers(),
		armredis.Caches(),
	}
}
