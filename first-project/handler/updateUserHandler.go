package handler

import (
	"net/http"

	"first-project/schemas"

	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(ctx *gin.Context) {
	request := UpdateUserRequest{}

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

	if err := DATABASE.Save(&user).Error; err != nil {
		LOGGER.Errf("Error updating user: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, http.StatusOK, "user updated", user)
}
