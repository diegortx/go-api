package routes

import (
	"database/sql"
	"go-api/routes/product"
	"go-api/routes/user"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(server *gin.Engine, dbConnection *sql.DB) {

	product.RegisterProductRoutes(server, dbConnection)
	user.RegisterUserRoutes(server, dbConnection)

	server.GET("/", func(ctx *gin.Context) {
		ctx.File("./views/home.html")
	})
}
