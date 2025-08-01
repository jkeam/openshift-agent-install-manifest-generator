package utils

import (
	context "context"
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Representation of an operator channel
type OperatorChannel struct {
	CsvName          string   `json:"csvName"`
	DisplayName      string   `json:"displayName"`
	AdditionalImages []string `json:"additionalImages"`
}

// Representation of an operator package
type OperatorPackage struct {
	PackageName        string                      `json:"packageName"`
	DefaultChannel     string                      `json:"defaultChannel"`
	DefaultDisplayName string                      `json:"defaultDisplayName"`
	Channels           map[string]*OperatorChannel `json:"channels"`
}

// OpenShift Registry Client
type OpenShiftRegistryClientInterface interface {
	ListPackages(ctx context.Context, in *ListPackageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PackageName], error)
	GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*Package, error)
	GetBundle(ctx context.Context, in *GetBundleRequest, opts ...grpc.CallOption) (*Bundle, error)
}
type OpenShiftRegistryClient struct {
	RegistryClient RegistryClient
}

func (o *OpenShiftRegistryClient) ListPackages(ctx context.Context, in *ListPackageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PackageName], error) {
	return o.RegistryClient.ListPackages(ctx, in, opts...)
}
func (o *OpenShiftRegistryClient) GetPackage(ctx context.Context, in *GetPackageRequest, opts ...grpc.CallOption) (*Package, error) {
	return o.RegistryClient.GetPackage(ctx, in, opts...)
}
func (o *OpenShiftRegistryClient) GetBundle(ctx context.Context, in *GetBundleRequest, opts ...grpc.CallOption) (*Bundle, error) {
	return o.RegistryClient.GetBundle(ctx, in, opts...)
}

// Returns new red hat operator marketplace client
func NewOpenShiftRegistryClient() *OpenShiftRegistryClient {
	// open connection
	conn, err := grpc.NewClient(
		"redhat-operators.openshift-marketplace.svc.cluster.local:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed connecting: %v", err)
	}
	defer conn.Close()

	// create client
	return &OpenShiftRegistryClient{RegistryClient: NewRegistryClient(conn)}
}
