package controllers

import (
	docs "github.com/NIDNetwork/nid-native-registry/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.BasePath = "/api/v1"
}

// SwaggerDocs .
var SwaggerDocs = ginSwagger.WrapHandler(swaggerfiles.Handler)
