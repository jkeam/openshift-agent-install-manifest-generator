package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"strings"
)

// Get Package Test
func TestGetPackages(t *testing.T) {
	msg := `{success: true}`
	router := setupRouter()
	router = getPackages(router, func () any {
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
	router := setupRouter()
	router = getPackageByName(router, func (name string) any {
		return name
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/packages/packageName", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, msg, strings.Trim(w.Body.String(), `"`))
}
