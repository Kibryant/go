package router

import (
	"first-project/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	r := router.Group("/api/v1")

	{
		r.GET("/users", handler.GetUsersHandler)

		r.GET("/user", handler.GetUserHandler)

		r.POST("/user", handler.CreateUserHandler)

		r.DELETE("/user", handler.DeleteUserHandler)

		r.PUT("/user", handler.UpdateUserHandler)
	}
}
