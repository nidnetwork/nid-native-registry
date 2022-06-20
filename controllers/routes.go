package controllers

import (
	"github.com/gin-gonic/gin"
)

// CreateRouter -
func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/swagger/*any", SwaggerDocs)
	r.GET("/", Index)
	r.GET("/api/v1", Index)

	return r
}
