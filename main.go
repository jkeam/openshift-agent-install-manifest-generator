package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jkeam/openshift-agent-install-manifest-generator/utils"
)

// Setup the router
func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowWildcard = true
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))
	return router
}

// Add getPackages endpoint
func getPackages(router *gin.Engine, client utils.OpenShiftRegistryClientInterface, handler func(utils.OpenShiftRegistryClientInterface) any) *gin.Engine {
	router.GET("/packages", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, handler(client))
	})
	return router
}

// Add getPackagesByName endpoint
func getPackageByName(router *gin.Engine, client utils.OpenShiftRegistryClientInterface, handler func(utils.OpenShiftRegistryClientInterface, string) any) *gin.Engine {
	router.GET("/packages/:packageName", func(c *gin.Context) {
		packageName := c.Param("packageName")
		c.IndentedJSON(http.StatusOK, handler(client, packageName))
	})
	return router
}

// Entrypoint
func main() {
	client := utils.NewOpenShiftRegistryClient()
	router := setupRouter()
	router = getPackages(router, client, utils.GetPackages)
	router = getPackageByName(router, client, utils.GetPackageByName)
	router.Run(":8080")
}
