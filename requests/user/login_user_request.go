package requests

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,max=250"`
}

func (r *LoginUserRequest) Validate() map[string]string {
	err := validate.Struct(r)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		var errorMsg string
		switch err.Field() {
		case "Email":
			if err.Tag() == "required" {
				errorMsg = "Email is required"
			} else if err.Tag() == "email" {
				errorMsg = "Email is not valid"
			}
		case "Password":
			if err.Tag() == "required" {
				errorMsg = "Password is required"
			} else if err.Tag() == "max" {
				errorMsg = "Password max length is 250"
			}
		default:
			errorMsg = fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", err.Field(), err.Tag())
		}
		errors[err.Field()] = errorMsg
	}
	return errors
}
