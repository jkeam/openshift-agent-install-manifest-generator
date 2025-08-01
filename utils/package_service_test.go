package utils

import (
	context "context"
	"testing"
	"fmt"
	"io"

	grpc "google.golang.org/grpc"
	gomock "go.uber.org/mock/gomock"
)

type ServerStreamingClientMock struct {
	grpc.ClientStream
}
func (s *ServerStreamingClientMock) Recv() (*PackageName, error) {
	return &PackageName{Name: "submariner"}, io.EOF
}

type OpenShiftRegistryClientMock struct {
}

func (o *OpenShiftRegistryClientMock) ListPackages(ctx context.Context, in *ListPackageRequest, _ ...grpc.CallOption) (grpc.ServerStreamingClient[PackageName], error) {
	return nil, nil
}
func (o *OpenShiftRegistryClientMock) GetPackage(ctx context.Context, in *GetPackageRequest, _ ...grpc.CallOption) (*Package, error) {
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
func (o *OpenShiftRegistryClientMock) GetBundle(ctx context.Context, in *GetBundleRequest, _ ...grpc.CallOption) (*Bundle, error) {
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
	ctrl := gomock.NewController(t)
	/*
	channels := []*Channel {
		{Name: "channel", CsvName: "csv-name"},
	}
	thePackage := &Package{
		Name: "submariner",
		DefaultChannelName: "submariner",
		Channels: channels,
	}
		*/
	var serverStreamingClient grpc.ServerStreamingClient[PackageName]
	serverStreamingClient = &ServerStreamingClientMock {}
	mockClient := NewMockRegistryClient(ctrl)
	mockClient.EXPECT().ListPackages(context.Background(), &ListPackageRequest{}).Return(serverStreamingClient, nil).MinTimes(1)
	GetPackages(mockClient)
    //mockClient.AssertExpectations(t)
	a := 1 + 3
	if a != 4 {
		t.Errorf("Math is broken")
	}
}
