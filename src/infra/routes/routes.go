package routes

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	userRouters(router.Group("/user"))

	return router
}
