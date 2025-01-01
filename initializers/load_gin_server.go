package initializers

import (
	"go-api/db"
	"go-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func LoadGinServer() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	routes.RegisterRoutes(server, dbConnection)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}
	server.Run(":" + port)
}
