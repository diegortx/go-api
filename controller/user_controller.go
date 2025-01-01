package controller

import (
	"go-api/model"
	requests "go-api/requests/user"
	"go-api/useCase"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	UserUseCase useCase.UserUseCase
}

func NewUserController(usecase useCase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: usecase,
	}
}

func (u *UserController) GetUsers(ctx *gin.Context) {
	users, err := u.UserUseCase.GetUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to get users",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    users,
	})
}

func (u *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "ID is required",
		})
		return
	}

	idNumeric, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{
			Message: "Invalid ID, can be number only",
			Data:    err.Error(),
		})
		return
	}

	user, err := u.UserUseCase.GetUserByID(idNumeric)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to get user",
			Data:    err.Error(),
		})
		return
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, model.Response{
			Message: "User not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Data:    user,
	})
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req requests.CreateUserRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userCreated, err := u.UserUseCase.CreateUser(user)

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "email already exists" {
			status = http.StatusConflict
		}

		ctx.JSON(status, model.Response{
			Message: "Failed to create user",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "Success",
		Data:    userCreated,
	})
}

func (u *UserController) Login(ctx *gin.Context) {
	var req requests.LoginUserRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, model.Response{
			Message: "Invalid request body",
			Data:    err.Error(),
		})
		return
	}

	user, err := u.UserUseCase.Login(req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, model.Response{
			Message: "Failed to login",
			Data:    "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{
			Message: "Failed to create token, try again",
			Data:    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Token:   tokenString,
	})
}
