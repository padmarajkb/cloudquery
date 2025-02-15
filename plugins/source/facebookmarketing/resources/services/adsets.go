package services

import (
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/client"
	"github.com/cloudquery/cloudquery/plugins/source/facebookmarketing/rest"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Adsets() *schema.Table {
	return &schema.Table{
		Name:        "facebookmarketing_adsets",
		Resolver:    fetchAdsets,
		Transform:   transformers.TransformWithStruct(&rest.Adset{}, append(client.TransformerOptions(), transformers.WithPrimaryKeys("Id"))...),
		Description: "https://developers.facebook.com/docs/marketing-api/reference/ad-campaign#Reading",
	}
}
