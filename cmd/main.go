package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/useCase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//Camada Repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//Camada UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	//Camada de Controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProduct)

	server.Run(":8080")
}
