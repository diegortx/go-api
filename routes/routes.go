package routes

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, dbConnection *sql.DB) {

	// Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Camada de Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.PUT("/product/:productId", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProductById)
}
