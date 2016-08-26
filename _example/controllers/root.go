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
		"companies_url": baseURL + "/api/companies",
		"company_url":   baseURL + "/api/companies/{id}",
		"emails_url":    baseURL + "/api/emails",
		"email_url":     baseURL + "/api/emails/{id}",
		"jobs_url":      baseURL + "/api/jobs",
		"job_url":       baseURL + "/api/jobs/{id}",
		"profiles_url":  baseURL + "/api/profiles",
		"profile_url":   baseURL + "/api/profiles/{id}",
		"users_url":     baseURL + "/api/users",
		"user_url":      baseURL + "/api/users/{id}",
	}

	c.IndentedJSON(http.StatusOK, resources)
}
