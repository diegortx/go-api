package product

import (
	"database/sql"
	"go-api/controller"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterProductRoutes registers the product routes
func RegisterProductRoutes(server *gin.Engine, dbConnection *sql.DB) {
	productController := initializeProductController(dbConnection)

	server.GET("/products", productController.GetProducts)
	server.POST("/product", productController.CreateProduct)
	server.GET("/product/:productId", productController.GetProductById)
	server.PUT("/product/:productId", productController.UpdateProduct)
	server.DELETE("/product/:productId", productController.DeleteProductById)
}

func initializeProductController(dbConnection *sql.DB) *controller.ProductController {
	// Camada Repository
	productRepository := repository.NewProductRepository(dbConnection)

	// Camada UseCase
	productUseCase := usecase.NewProductUseCase(productRepository)

	// Camada de Controllers
	productController := controller.NewProductController(productUseCase)

	return productController
}
