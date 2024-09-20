package handler

import (
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{"error": message})
}

func sendSuccess(ctx *gin.Context, code int, message string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"data":    data,
	})
}
