package handler

import (
	"first-project/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := DATABASE.Find(&users).Error; err != nil {
		LOGGER.Errf("Error getting users: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error getting users")
		return
	}

	sendSuccess(ctx, http.StatusOK, "users found", users)
}
