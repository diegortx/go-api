package user

import (
	"database/sql"
	"go-api/controller"
	"go-api/middleware"
	"go-api/repository"
	"go-api/useCase"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine, dbConnection *sql.DB) {
	userController := initializeUserController(dbConnection)

	server.POST("/api/v1/login", userController.Login)

	prefix := server.Group("/api/v1/users")

	prefix.Use(middleware.RequireAuth(dbConnection))

	prefix.GET("/", userController.GetUsers)
	prefix.GET("/:id", userController.GetUserByID)
	prefix.POST("/", userController.CreateUser)
}

func initializeUserController(dbConnection *sql.DB) *controller.UserController {
	// Camada Repository
	userRepository := repository.NewUserRepository(dbConnection)

	// Camada UseCase
	userUseCase := useCase.NewUserUseCase(userRepository)

	// Camada de Controllers
	userController := controller.NewUserController(userUseCase)

	return userController
}
