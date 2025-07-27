package main

import (
  "context"
  "log"
  "io"
  "net/http"
  "encoding/json"

  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "google.golang.org/grpc"
  "github.com/jkeam/openshift-agent-install-manifest-generator/pkg/api"
  v1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
)

func getPackageNames(c *gin.Context) {
  conn, err := grpc.NewClient("redhat-operators.openshift-marketplace.svc.cluster.local:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  client := api.NewRegistryClient(conn)

  // fetch all the published operators
  resp, err := client.ListPackages(context.Background(), &api.ListPackageRequest{})
  var packages []*GetPackageResponse
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }
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

  c.IndentedJSON(http.StatusOK, packages) 
}

type ChannelResponse struct {
  CsvName           string    `json:"csvName"`
  DisplayName       string    `json:"displayName"`
  AdditionalImages  []string  `json:"additionalImages"`
}

type GetPackageResponse struct {
  PackageName         string            `json:"packageName"`
  DefaultChannel      string            `json:"defaultChannel"`
  DefaultDisplayName  string            `json:"defaultDisplayName"`
  Channels      map[string]*ChannelResponse `json:"channels"`
}

func getPackageByName(packageName string) *GetPackageResponse {
  httpResp := &GetPackageResponse{Channels: make(map[string]*ChannelResponse)}

  conn, err := grpc.NewClient("redhat-operators.openshift-marketplace.svc.cluster.local:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()

  client := api.NewRegistryClient(conn)

  resp, err := client.GetPackage(context.Background(), &api.GetPackageRequest{Name: packageName})
  if err != nil {
    log.Fatalf("Could not get package: %v", err)
  }
  
  httpResp.PackageName = resp.GetName()
  httpResp.DefaultChannel = resp.GetDefaultChannelName()
  
  for _, element := range resp.GetChannels() {
    channelResponse := &ChannelResponse{CsvName: element.GetCsvName()}
    
    // Get bundle for channel
    channelResp, channelErr := client.GetBundle(context.Background(), &api.GetBundleRequest{PkgName: resp.GetName(), ChannelName: element.GetName(), CsvName: element.GetCsvName()})
    if channelErr != nil {
      log.Fatalf("Could not get channel: %v", channelErr)
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

func getPackage(c *gin.Context) {
  packageName := c.Param("packageName")
  
  httpResp := getPackageByName(packageName)

  c.IndentedJSON(http.StatusOK, httpResp)
}

func main() {
  router := gin.Default()

  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"*"}
  config.AllowWildcard = true
  config.AllowHeaders = []string{"Content-Type"}

  router.Use(cors.New(config))
  router.GET("/packages", getPackageNames)
  router.GET("/packages/:packageName", getPackage)

  router.Run("localhost:8080")
}
