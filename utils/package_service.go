package utils

import (
	context "context"
	"encoding/json"
	"io"
	"log"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
)

// HELPER FUNCTIONS

// Helper function to get a specific package by name
func getPackageByName(client *OpenShiftRegistryClient, packageName string) *OperatorPackage {
	thePackage, err := client.RegistryClient.GetPackage(context.Background(), &GetPackageRequest{Name: packageName})
	if err != nil {
		log.Fatalf("Failed getting package: %v", err)
	}

	// process results
	operatorPackage := &OperatorPackage{Channels: make(map[string]*OperatorChannel)}
	operatorPackage.PackageName = thePackage.GetName()
	operatorPackage.DefaultChannel = thePackage.GetDefaultChannelName()

	for _, element := range thePackage.GetChannels() {
		channel := &OperatorChannel{CsvName: element.GetCsvName()}

		// Get bundle for channel
		bundle, bundleErr := client.RegistryClient.GetBundle(
			context.Background(),
			&GetBundleRequest{
				PkgName:     thePackage.GetName(),
				ChannelName: element.GetName(),
				CsvName:     element.GetCsvName(),
			},
		)
		if bundleErr != nil {
			log.Fatalf("Failed getting bundle for channel: %v", bundleErr)
		}

		// unmarshall into channel response
		var csv v1alpha1.ClusterServiceVersion
		csvBytes := []byte(bundle.GetCsvJson())
		json.Unmarshal(csvBytes, &csv)

		// set name and populate the channel additional images
		channel.DisplayName = csv.Spec.DisplayName
		for _, additionalImage := range csv.Spec.RelatedImages {
			// TODO: Check to see if Image is already in AdditionalImages
			channel.AdditionalImages = append(channel.AdditionalImages, additionalImage.Image)
		}
		operatorPackage.Channels[element.GetName()] = channel
	}

	operatorPackage.DefaultDisplayName = operatorPackage.Channels[operatorPackage.DefaultChannel].DisplayName
	return operatorPackage
}

func getPackages(client *OpenShiftRegistryClient) []*OperatorPackage {
	packageListing, err := client.RegistryClient.ListPackages(context.Background(), &ListPackageRequest{})
	if err != nil {
		log.Fatalf("Failed listing packages: %v", err)
	}

	// process output
	var packages []*OperatorPackage
	for {
		thePackage, err := packageListing.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListPackages(_) = _, %v", client, err)
		}
		packageName := thePackage.GetName()
		// TODO: make sure we haven't processed packageName already
		packages = append(packages, getPackageByName(client, packageName))
	}
	return packages
}

// EXTERNAL FUNCTIONS

// Get operator packages
func GetPackages(client *OpenShiftRegistryClient) any {
	return getPackages(client)
}

// Get package by name
func GetPackageByName(client *OpenShiftRegistryClient, name string) any {
	return getPackageByName(client, name)
}
