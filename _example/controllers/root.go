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
		"companies_url": baseURL + "/companies",
		"company_url":   baseURL + "/companies/{id}",
		"emails_url":    baseURL + "/emails",
		"email_url":     baseURL + "/emails/{id}",
		"jobs_url":      baseURL + "/jobs",
		"job_url":       baseURL + "/jobs/{id}",
		"profiles_url":  baseURL + "/profiles",
		"profile_url":   baseURL + "/profiles/{id}",
		"users_url":     baseURL + "/users",
		"user_url":      baseURL + "/users/{id}",
	}

	c.IndentedJSON(http.StatusOK, resources)
}
