package utils

import (
	context "context"
	"fmt"
	"io"
	"testing"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
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
	channels := []*Channel{
		{Name: "channel", CsvName: "csv-name"},
	}
	thePackage := &Package{
		Name:               in.Name,
		DefaultChannelName: in.Name,
		Channels:           channels,
	}

	return thePackage, nil
}
func (o *OpenShiftRegistryClientMock) GetBundle(ctx context.Context, in *GetBundleRequest, _ ...grpc.CallOption) (*Bundle, error) {
	bundle := &Bundle{
		CsvJson: fmt.Sprintf(`{"spec": {displayName: "%s", "relatedImages": []}}`, in.PkgName),
	}
	return bundle, nil
}

func TestGetPackageByName(t *testing.T) {
	packageName := "submariner"
	channelName := "channel"
	csvName := "csv-name"
	// get package
	channels := []*Channel{
		{Name: channelName, CsvName: csvName},
	}
	thePackage := &Package{
		Name:               packageName,
		DefaultChannelName: packageName,
		Channels:           channels,
	}
	getPackageRequest := &GetPackageRequest{
		Name: packageName,
	}

	// get bundle
	bundle := &Bundle{
		CsvJson: fmt.Sprintf(`{"spec": {displayName: "%s", "relatedImages": []}}`, packageName),
	}
	getBundleRequest := &GetBundleRequest {
		PkgName: packageName,
		ChannelName: channelName,
		CsvName: csvName,
	}

	ctrl := gomock.NewController(t)
	mockClient := NewMockRegistryClient(ctrl)
	mockClient.EXPECT().GetPackage(context.Background(), getPackageRequest).Return(thePackage, nil).MinTimes(1)
	mockClient.EXPECT().GetBundle(context.Background(), getBundleRequest).Return(bundle, nil).MinTimes(1)
	GetPackageByName(mockClient, packageName)
}

func TestGetPackages(t *testing.T) {
	ctrl := gomock.NewController(t)
	var serverStreamingClient grpc.ServerStreamingClient[PackageName]
	serverStreamingClient = &ServerStreamingClientMock{}
	mockClient := NewMockRegistryClient(ctrl)
	mockClient.EXPECT().ListPackages(context.Background(), &ListPackageRequest{}).Return(serverStreamingClient, nil).MinTimes(1)
	GetPackages(mockClient)
}
