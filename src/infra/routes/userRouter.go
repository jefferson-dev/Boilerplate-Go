package routes

import (
	"github.com/gin-gonic/gin"

	userController "boilerplate/src/modules/user/controller"
)

func userRouters(router *gin.RouterGroup) {
	router.POST("/", userController.CreateUser)
	router.GET("/:id", userController.GetOneUser)
	router.GET("/", userController.GetAllUser)
	router.PUT("/:id", userController.UpdateUser)
	router.DELETE("/:id", userController.DeleteUser)
}
