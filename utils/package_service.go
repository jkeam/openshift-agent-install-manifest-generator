package utils

import (
	context "context"
	"encoding/json"
	"io"
	"log"

	"github.com/operator-framework/api/pkg/operators/v1alpha1"
)

// HELPER FUNCTIONS
func appendUnique(tracker map[string]bool, list []string, value string) []string {
	if tracker[value] {
		return list
	}
	list = append(list, value)
	tracker[value] = true
	return list
}

// Helper function to get a specific package by name
func getPackageByName(client OpenShiftRegistryClientInterface, packageName string) *OperatorPackage {
	thePackage, err := client.GetPackage(context.Background(), &GetPackageRequest{Name: packageName})
	if err != nil || thePackage == nil {
		log.Fatalf("Failed getting package: %v", err)
	}

	// process results
	operatorPackage := &OperatorPackage{Channels: make(map[string]*OperatorChannel)}
	operatorPackage.PackageName = thePackage.GetName()
	operatorPackage.DefaultChannel = thePackage.GetDefaultChannelName()

	for _, element := range thePackage.GetChannels() {
		channel := &OperatorChannel{CsvName: element.GetCsvName()}

		// Get bundle for channel
		bundle, bundleErr := client.GetBundle(
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
		tracker := make(map[string]bool)
		for _, additionalImage := range csv.Spec.RelatedImages {
			channel.AdditionalImages = appendUnique(tracker, channel.AdditionalImages, additionalImage.Image)
		}
		operatorPackage.Channels[element.GetName()] = channel
	}

	defaultChannel := operatorPackage.Channels[operatorPackage.DefaultChannel]
	if defaultChannel != nil {
		operatorPackage.DefaultDisplayName = defaultChannel.DisplayName
	}
	return operatorPackage
}

func getPackages(client OpenShiftRegistryClientInterface) []*OperatorPackage {
	packageListing, err := client.ListPackages(context.Background(), &ListPackageRequest{})
	if err != nil || packageListing == nil {
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
		packages = append(packages, getPackageByName(client, packageName))
	}
	return packages
}

// EXTERNAL FUNCTIONS

// Get operator packages
func GetPackages(client OpenShiftRegistryClientInterface) any {
	return getPackages(client)
}

// Get package by name
func GetPackageByName(client OpenShiftRegistryClientInterface, name string) any {
	return getPackageByName(client, name)
}
