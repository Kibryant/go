package handler

import (
	"first-project/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParamater").Error())
		return
	}

	user := schemas.User{}

	if err := DATABASE.Where("id = ?", id).First(&user).Error; err != nil {
		LOGGER.Errf("Error getting user: %s", err.Error())
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	sendSuccess(ctx, http.StatusOK, "user found", user)
}
