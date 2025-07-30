package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jkeam/openshift-agent-install-manifest-generator/utils"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowWildcard = true
	config.AllowHeaders = []string{"Content-Type"}

	router.Use(cors.New(config))
	router.GET("/packages", utils.GetPackagesRoute)
	router.GET("/packages/:packageName", utils.GetPackageByNameRoute)

	router.Run(":8080")
}
