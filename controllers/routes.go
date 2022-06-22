package controllers

import (
	"github.com/gin-gonic/gin"
)

// CreateRouter -
func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/docs/*any", SwaggerDocs)
	r.GET("/", Index)
	r.GET("/api/v1", Index)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/records/:id", GetRecord)
		v1.POST("/records", CreateRecord)
		v1.PATCH("/records/:id", UpdateRecord)
		v1.DELETE("/records/:id", DeleteRecord)
	}

	return r
}
