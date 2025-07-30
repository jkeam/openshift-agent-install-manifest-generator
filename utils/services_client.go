package utils

import (
	"log"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Returns new red hat operator marketplace client
func NewOpenShiftRegistryClient() RegistryClient {
	// open connection
	conn, err := grpc.NewClient(
		"redhat-operators.openshift-marketplace.svc.cluster.local:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed connecting: %v", err)
	}
	defer conn.Close()

	// create client
	return NewRegistryClient(conn)
}
