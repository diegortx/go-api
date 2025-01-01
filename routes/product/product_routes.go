package product

import (
	"database/sql"
	"go-api/controller"
	"go-api/middleware"
	"go-api/repository"
	"go-api/useCase"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers the product routes
func RegisterProductRoutes(server *gin.Engine, dbConnection *sql.DB) {
	productController := initializeProductController(dbConnection)

	prefix := server.Group("/api/v1/product")
	prefix.Use(middleware.RequireAuth(dbConnection))

	prefix.GET("/", productController.GetProducts)
	prefix.POST("/", productController.CreateProduct)
	prefix.GET("/:productId", productController.GetProductById)
	prefix.PUT("/:productId", productController.UpdateProduct)
	prefix.DELETE("/:productId", productController.DeleteProductById)
}

func initializeProductController(dbConnection *sql.DB) *controller.ProductController {
	// Camada Repository
	productRepository := repository.NewProductRepository(dbConnection)

	// Camada UseCase
	productUseCase := useCase.NewProductUseCase(productRepository)

	// Camada de Controllers
	productController := controller.NewProductController(productUseCase)

	return productController
}
