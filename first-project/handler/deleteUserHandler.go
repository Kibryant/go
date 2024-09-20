package handler

import (
	"first-project/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context) {
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

	if err := DATABASE.Delete(&user).Error; err != nil {
		LOGGER.Errf("Error deleting user: %s", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error deleting user")
		return
	}

	sendSuccess(ctx, http.StatusOK, "user deleted", nil)
}
