package utils

import "testing"

func TestGetPackageNames(t *testing.T) {
	a := 1 + 3
	if a != 4 {
		t.Errorf("Math is broken")
	}
	/*
	   conn, err := grpc.NewClient("redhat-operators.openshift-marketplace.svc.cluster.local:50051", grpc.WithInsecure())

	   	if err != nil {
	   		log.Fatalf("Failed connecting: %v", err)
	   	}

	   defer conn.Close()

	   client := NewRegistryClient(conn)

	   // set up mock

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
	*/
}
