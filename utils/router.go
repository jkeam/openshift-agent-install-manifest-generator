package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPackagesRoute(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, GetPackages())
}
func GetPackageByNameRoute(c *gin.Context) {
	packageName := c.Param("packageName")
	c.IndentedJSON(http.StatusOK, GetPackageByName(packageName))
}
