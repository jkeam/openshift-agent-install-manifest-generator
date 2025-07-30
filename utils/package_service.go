package utils

import (
	context "context"
	"encoding/json"
	"io"
	"log"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
)

// INTERNAL HELPER FUNCTIONS

// Helper function to get a specific package by name
func getPackageByName(packageName string) *OperatorPackage {
	// create client
	client := NewOpenShiftRegistryClient()
	resp, err := client.GetPackage(context.Background(), &GetPackageRequest{Name: packageName})
	if err != nil {
		log.Fatalf("Failed getting package: %v", err)
	}

	// process results
	httpResp := &OperatorPackage{Channels: make(map[string]*OperatorChannel)}
	httpResp.PackageName = resp.GetName()
	httpResp.DefaultChannel = resp.GetDefaultChannelName()

	for _, element := range resp.GetChannels() {
		channelResponse := &OperatorChannel{CsvName: element.GetCsvName()}

		// Get bundle for channel
		channelResp, channelErr := client.GetBundle(
			context.Background(),
			&GetBundleRequest{
				PkgName:     resp.GetName(),
				ChannelName: element.GetName(),
				CsvName:     element.GetCsvName(),
			},
		)
		if channelErr != nil {
			log.Fatalf("Failed getting bundle for channel: %v", channelErr)
		}

		// unmarshall into channel response
		var csv v1alpha1.ClusterServiceVersion
		csvBytes := []byte(channelResp.GetCsvJson())
		json.Unmarshal(csvBytes, &csv)

		// set name and populate the channel additional images
		channelResponse.DisplayName = csv.Spec.DisplayName
		for _, additionalImage := range csv.Spec.RelatedImages {
			// TODO: Check to see if Image is already in AdditionalImages
			channelResponse.AdditionalImages = append(channelResponse.AdditionalImages, additionalImage.Image)
		}
		httpResp.Channels[element.GetName()] = channelResponse
	}

	httpResp.DefaultDisplayName = httpResp.Channels[httpResp.DefaultChannel].DisplayName
	return httpResp
}

func getPackages() []*OperatorPackage {
	// create client
	client := NewOpenShiftRegistryClient()
	resp, err := client.ListPackages(context.Background(), &ListPackageRequest{})
	if err != nil {
		log.Fatalf("Failed listing packages: %v", err)
	}

	// process output
	var packages []*OperatorPackage
	for {
		packageName, err := resp.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPackages(_) = _, %v", client, err)
		}
		packages = append(packages, getPackageByName(packageName.GetName()))
	}
	return packages
}

// EXPORTED FUNCTIONS

// Get operator package names
func GetPackages() any {
	return getPackages()
}

// Get package by name
func GetPackageByName(name string) any {
	return getPackageByName(name)
}
