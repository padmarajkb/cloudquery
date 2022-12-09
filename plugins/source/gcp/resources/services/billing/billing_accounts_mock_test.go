// Code generated by codegen; DO NOT EDIT.

package billing

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"testing"

	pb "google.golang.org/genproto/googleapis/cloud/billing/v1"
)

func createBillingAccounts(gsrv *grpc.Server) error {
	fakeServer := &fakeBillingAccountsServer{}
	pb.RegisterCloudBillingServer(gsrv, fakeServer)
	return nil
}

type fakeBillingAccountsServer struct {
	pb.UnimplementedCloudBillingServer
}

func (f *fakeBillingAccountsServer) ListBillingAccounts(context.Context, *pb.ListBillingAccountsRequest) (*pb.ListBillingAccountsResponse, error) {
	resp := pb.ListBillingAccountsResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestBillingAccounts(t *testing.T) {
	client.MockTestGrpcHelper(t, BillingAccounts(), createBillingAccounts, client.TestOptions{})
}
