package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func OfflineConversionDataSets() *schema.Table {
	return &schema.Table{
		Name:      "facebookmarketing_offline_conversion_data_sets",
		Resolver:  fetchOfflineConversionDataSets,
		Transform: transformers.TransformWithStruct(&rest.OfflineConversionDataSet{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Resolver: client.ResolveAccountId,
				Type:     schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Description: "https://developers.facebook.com/docs/marketing-api/reference/offline-conversion-data-set/#Reading",
	}
}
