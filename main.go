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
func getPackages(router *gin.Engine) *gin.Engine {
	router.GET("/packages", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, utils.GetPackages())
	})
	return router
}

// Add getPackagesByName endpoint
func getPackageByName(router *gin.Engine) *gin.Engine {
	router.GET("/packages/:packageName", func(c *gin.Context) {
		packageName := c.Param("packageName")
		c.IndentedJSON(http.StatusOK, utils.GetPackageByName(packageName))
	})
	return router
}

// Entrypoint
func main() {
	router := setupRouter()
	router = getPackages(router)
	router = getPackageByName(router)
	router.Run(":8080")
}
