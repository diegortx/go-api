package main

import (
	"go-api/db"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	routes.RegisterRoutes(server, dbConnection)

	server.Run(":8080")
}
