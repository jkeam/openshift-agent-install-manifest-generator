package utils

import (
	"testing"
	grpc "google.golang.org/grpc"
)

func TestGetPackageNames(t *testing.T) {
	conn, err := grpc.NewClient("redhat-operators.openshift-marketplace.svc.cluster.local:50051", grpc.WithInsecure())
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
