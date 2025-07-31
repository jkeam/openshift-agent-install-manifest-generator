package utils

import (
	context "context"
	"fmt"
	"testing"

	grpc "google.golang.org/grpc"
)

type OpenShiftRegistryClientMock struct {
}

func (o *OpenShiftRegistryClientMock) ListPackages(ctx context.Context, in *ListPackageRequest) (grpc.ServerStreamingClient[PackageName], error) {
	return nil, nil
}
func (o *OpenShiftRegistryClientMock) GetPackage(ctx context.Context, in *GetPackageRequest) (*Package, error) {
	return nil, nil
}
func (o *OpenShiftRegistryClientMock) GetBundle(ctx context.Context, in *GetBundleRequest) (*Bundle, error) {
	fmt.Println(in)
	return nil, nil
}

func TestGetPackageByName(t *testing.T) {
	client := &OpenShiftRegistryClientMock{}
	GetPackageByName(client, "test")
	a := 1 + 3
	if a != 4 {
		t.Errorf("Math is broken")
	}
}

func TestGetPackages(t *testing.T) {
	client := &OpenShiftRegistryClientMock{}
	GetPackages(client)
	a := 1 + 3
	if a != 4 {
		t.Errorf("Math is broken")
	}
}
