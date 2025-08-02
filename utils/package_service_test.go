package utils

import (
	context "context"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

type ServerStreamingClientMock struct {
	grpc.ClientStream
	RecvCount int
}

func (s *ServerStreamingClientMock) Recv() (*PackageName, error) {
	s.RecvCount++
	return &PackageName{Name: "submariner"}, io.EOF
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
	getBundleRequest := &GetBundleRequest{
		PkgName:     packageName,
		ChannelName: channelName,
		CsvName:     csvName,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := NewMockRegistryClient(ctrl)
	mockClient.EXPECT().GetPackage(context.Background(), getPackageRequest).Return(thePackage, nil).MinTimes(1)
	mockClient.EXPECT().GetBundle(context.Background(), getBundleRequest).Return(bundle, nil).MinTimes(1)
	GetPackageByName(mockClient, packageName)
}

func TestGetPackages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serverStreamingClient := &ServerStreamingClientMock{}
	mockClient := NewMockRegistryClient(ctrl)
	mockClient.EXPECT().ListPackages(context.Background(), &ListPackageRequest{}).Return(serverStreamingClient, nil).MinTimes(1)
	GetPackages(mockClient)
	assert.Equal(t, 1, serverStreamingClient.RecvCount)
}

func TestAppendUnique(t *testing.T) {
	tracker := make(map[string]bool)
	fruits := []string{"apple", "banana", "cherry", "apple", "banana", "cherry", "cherry"}
	uniqueFruits := []string{"apple", "banana", "cherry"}
	actualList := []string{}
	for _, fruit := range fruits {
		actualList = appendUnique(tracker, actualList, fruit)
	}
	assert.Equal(t, uniqueFruits, actualList)
}
