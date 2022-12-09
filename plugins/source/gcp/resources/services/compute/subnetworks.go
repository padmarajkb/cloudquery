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

func Subnetworks() *schema.Table {
	return &schema.Table{
		Name:      "gcp_compute_subnetworks",
		Resolver:  fetchSubnetworks,
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
				Name:     "enable_flow_logs",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnableFlowLogs"),
			},
			{
				Name:     "external_ipv6_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ExternalIpv6Prefix"),
			},
			{
				Name:     "fingerprint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Fingerprint"),
			},
			{
				Name:     "gateway_address",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GatewayAddress"),
			},
			{
				Name:     "id",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "internal_ipv6_prefix",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InternalIpv6Prefix"),
			},
			{
				Name:     "ip_cidr_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IpCidrRange"),
			},
			{
				Name:     "ipv6_access_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ipv6AccessType"),
			},
			{
				Name:     "ipv6_cidr_range",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Ipv6CidrRange"),
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "log_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LogConfig"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "network",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Network"),
			},
			{
				Name:     "private_ip_google_access",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("PrivateIpGoogleAccess"),
			},
			{
				Name:     "private_ipv6_google_access",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PrivateIpv6GoogleAccess"),
			},
			{
				Name:     "purpose",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Purpose"),
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Region"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "secondary_ip_ranges",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecondaryIpRanges"),
			},
			{
				Name:     "stack_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StackType"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
		},
	}
}

func fetchSubnetworks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	req := &pb.AggregatedListSubnetworksRequest{
		Project: c.ProjectId,
	}
	gcpClient, err := compute.NewSubnetworksRESTClient(ctx, c.ClientOptions...)
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

		res <- resp.Value.Subnetworks

	}
	return nil
}
