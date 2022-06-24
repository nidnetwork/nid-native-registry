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
		v1.POST("/records", CreateRecord)
		v1.GET("/records/:nns", GetRecord)
		v1.PATCH("/records/:nns", UpdateRecord)
		v1.DELETE("/records/:nns", DeleteRecord)
	}

	return r
}
