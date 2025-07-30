package utils

import (
	context "context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/operator-framework/api/pkg/operators/v1alpha1"
)

func GetPackageNames(c *gin.Context) {
	conn, err := NewMarketPlaceClient()
	if err != nil {
		log.Fatalf("Failed connecting: %v", err)
	}
	defer conn.Close()

	client := NewRegistryClient(conn)

	// fetch all the published operators
	resp, err := client.ListPackages(context.Background(), &ListPackageRequest{})
	var packages []*GetPackageResponse
	if err != nil {
		log.Fatalf("Failed listing packages: %v", err)
	}
	for {
		packageName, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPackages(_) = _, %v", client, err)
		}
		packages = append(packages, GetPackageByName(packageName.GetName()))
	}

	c.IndentedJSON(http.StatusOK, packages)
}

func GetPackageByName(packageName string) *GetPackageResponse {
	conn, err := NewMarketPlaceClient()
	if err != nil {
		log.Fatalf("Failed connecting: %v", err)
	}
	defer conn.Close()

	client := NewRegistryClient(conn)

	resp, err := client.GetPackage(context.Background(), &GetPackageRequest{Name: packageName})
	if err != nil {
		log.Fatalf("Failed getting package: %v", err)
	}

	httpResp := &GetPackageResponse{Channels: make(map[string]*ChannelResponse)}
	httpResp.PackageName = resp.GetName()
	httpResp.DefaultChannel = resp.GetDefaultChannelName()

	for _, element := range resp.GetChannels() {
		channelResponse := &ChannelResponse{CsvName: element.GetCsvName()}

		// Get bundle for channel
		channelResp, channelErr := client.GetBundle(context.Background(), &GetBundleRequest{PkgName: resp.GetName(), ChannelName: element.GetName(), CsvName: element.GetCsvName()})
		if channelErr != nil {
			log.Fatalf("Failed getting bundle for channel: %v", channelErr)
		}

		var csv v1alpha1.ClusterServiceVersion
		json.Unmarshal([]byte(channelResp.GetCsvJson()), &csv)

		channelResponse.DisplayName = csv.Spec.DisplayName

		// populate the channel additional images
		for _, additionalImage := range csv.Spec.RelatedImages {
			channelResponse.AdditionalImages = append(channelResponse.AdditionalImages, additionalImage.Image)
		}

		httpResp.Channels[element.GetName()] = channelResponse
	}

	httpResp.DefaultDisplayName = httpResp.Channels[httpResp.DefaultChannel].DisplayName

	return httpResp
}

func GetPackage(c *gin.Context) {
	packageName := c.Param("packageName")

	httpResp := GetPackageByName(packageName)

	c.IndentedJSON(http.StatusOK, httpResp)
}
