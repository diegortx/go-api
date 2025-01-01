package middleware

import (
	"database/sql"
	"fmt"
	"go-api/controller"
	"go-api/model"
	"go-api/repository"
	"go-api/useCase"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(dbConnection *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//get token in header
		tokenString := ctx.GetHeader("token")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, model.Response{
				Message: "Not authorized",
				Data:    "Token is required",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, model.Response{
				Message: "Not authorized",
				Data:    err.Error(),
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {

			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				ctx.JSON(http.StatusUnauthorized, model.Response{
					Message: "Not authorized",
					Data:    "Token is expired",
				})
			}

			userID := int(claims["sub"].(float64))

			userRepository := repository.NewUserRepository(dbConnection)
			userUseCase := useCase.NewUserUseCase(userRepository)
			userController := controller.NewUserController(userUseCase)

			user, err := userController.UserUseCase.GetUserByID(userID)

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, model.Response{
					Message: "Not authorized",
					Data:    "User not found",
				})
				ctx.Abort()
				return
			}
			ctx.Set("user", user)
			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, model.Response{
				Message: "Not authorized",
				Data:    "Token is expired or invalid",
			})
			ctx.Abort()
			return
		}
	}
}
