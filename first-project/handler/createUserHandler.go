package handler

import (
	"first-project/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		LOGGER.Errf("Validation error: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Salary:   request.Salary,
	}

	if err := DATABASE.Create(&user).Error; err != nil {
		LOGGER.Errf("Error creating user: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user")
		return
	}

	sendSuccess(ctx, http.StatusCreated, "user created", user)
}
