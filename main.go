package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jkeam/openshift-agent-install-manifest-generator/utils"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowWildcard = true
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))
	return router
}

func getPackages(router *gin.Engine) *gin.Engine {
	router.GET("/packages", utils.GetPackagesRoute)
	return router
}

func getPackageByName(router *gin.Engine) *gin.Engine {
	router.GET("/packages/:packageName", utils.GetPackageByNameRoute)
	return router
}

func main() {
	router := setupRouter()
	router = getPackages(router)
	router = getPackageByName(router)
	router.Run(":8080")
}
