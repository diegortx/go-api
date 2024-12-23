package routes

import (
	"database/sql"
	"go-api/routes/product"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(server *gin.Engine, dbConnection *sql.DB) {

	product.RegisterProductRoutes(server, dbConnection)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
