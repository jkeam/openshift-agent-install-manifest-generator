package utils

import grpc "google.golang.org/grpc"

// Returns new red hat operator marketplace client
func NewMarketPlaceClient() (conn *grpc.ClientConn, err error) {
	return grpc.NewClient("redhat-operators.openshift-marketplace.svc.cluster.local:50051", grpc.WithInsecure())
}
