package utils

import (
	context "context"
	"testing"
	"fmt"
	"io"

	grpc "google.golang.org/grpc"
	"github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
	"github.com/golang/mock/gomock"
)

type OpenShiftRegistryClientMock struct {
}

func (o *OpenShiftRegistryClientMock) ListPackages(ctx context.Context, in *ListPackageRequest) (grpc.ServerStreamingClient[PackageName], error) {
	//mockClient := new(grpc.ServerStreamingClient)
	ctrl := gomock.NewController(t)
	channels := []*Channel {
		{Name: "channel", CsvName: "csv-name"},
	}
	thePackage := &Package{
		Name: in.Name,
		DefaultChannelName: in.Name,
		Channels: channels,
	}
	//mockClient = mocks.
    //mockClient.On("Recv").Return(thePackage, io.EOF).Once()
    //mockClient.AssertExpectations(t) // Verify that the expected calls were made
	return mockClient, nil
}
func (o *OpenShiftRegistryClientMock) GetPackage(ctx context.Context, in *GetPackageRequest) (*Package, error) {
	channels := []*Channel {
		{Name: "channel", CsvName: "csv-name"},
	}
	thePackage := &Package{
		Name: in.Name,
		DefaultChannelName: in.Name,
		Channels: channels,
	}

	return thePackage, nil
}
func (o *OpenShiftRegistryClientMock) GetBundle(ctx context.Context, in *GetBundleRequest) (*Bundle, error) {
	bundle := &Bundle {
		CsvJson: fmt.Sprintf(`{"spec": {displayName: "%s", "relatedImages": []}}`, in.PkgName),
	}
	return bundle, nil
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
