package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/jkeam/openshift-agent-install-manifest-generator/utils"
	"github.com/stretchr/testify/assert"
)

// Get Package Test
func TestGetPackages(t *testing.T) {
	msg := `{success: true}`
	client := &utils.OpenShiftRegistryClient{}
	router := setupRouter()
	router = getPackages(router, client, func(c *utils.OpenShiftRegistryClient) any {
		return msg
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/packages", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, msg, strings.Trim(w.Body.String(), `"`))
}

// Get Package by Name
func TestGetPackageByName(t *testing.T) {
	msg := `packageName`
	client := &utils.OpenShiftRegistryClient{}
	router := setupRouter()
	router = getPackageByName(router, client, func(c *utils.OpenShiftRegistryClient, name string) any {
		return name
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/packages/packageName", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, msg, strings.Trim(w.Body.String(), `"`))
}
