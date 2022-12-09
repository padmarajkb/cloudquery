// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"google.golang.org/api/iterator"

	pb "google.golang.org/genproto/googleapis/cloud/compute/v1"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugins/source/gcp/client"

	"cloud.google.com/go/compute/apiv1"
)

func SslCertificates() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_ssl_certificates",
		Resolver:  fetchSslCertificates,
		Multiplex: client.ProjectMultiplex,
		Columns: []schema.Column{
			{
				Name:     "project_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveProject,
			},
			{
				Name: "self_link",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Certificate"),
			},
			{
				Name:     "creation_timestamp",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreationTimestamp"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "expire_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExpireTime"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "managed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Managed"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "private_key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateKey"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "self_managed",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfManaged"),
			},
			{
				Name:     "subject_alternative_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SubjectAlternativeNames"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
		},
	}
}

func fetchSslCertificates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListSslCertificatesRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewSslCertificatesRESTClient(ctx, c.ClientOptions...)
	if err != nil {
		return err
	}
	it := gcpClient.AggregatedList(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		res <- resp.Value.SslCertificates

	}
	return nil
}
