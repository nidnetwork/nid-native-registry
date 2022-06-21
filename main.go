package main

import (
	"os"

	"github.com/nidnetwork/nid-native-registry/controllers"
	"github.com/nidnetwork/nid-native-registry/models"
)

// @title NID Native Registry API
// @version 1.0
// @description NID NNS registry server.
// @termsOfService https://nid.network/terms/

// @contact.name API Support
// @contact.url https://nid.network/support/
// @contact.email support@nid.network

// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Connect to database
	models.ConnectDatabase()

	r := controllers.CreateRouter()

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "127.0.0.1:8081"
	}

	r.Run(port)
}
