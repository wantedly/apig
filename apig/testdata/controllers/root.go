package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIEndpoints(c *gin.Context) {
	var reqScheme string
	if c.Request.TLS != nil {
		reqScheme = "https"
	} else {
		reqScheme = "http"
	}
	reqHost := c.Request.Host
	baseURL := fmt.Sprintf("%s://%s", reqScheme, reqHost)

	resources := map[string]string{
		"users_url": baseURL + "/users",
		"user_url":  baseURL + "/users/{id}",
	}

	c.IndentedJSON(http.StatusOK, resources)
}
