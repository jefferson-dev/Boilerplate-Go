package userController

import (
	"boilerplate/src/modules/user/domain/entity"
	"boilerplate/src/modules/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := entity.User{}

	c.BindJSON(&user)
	services.CreateUser(&user)
	c.JSON(http.StatusCreated, user)
}

func GetAllUser(c *gin.Context) {
	user := services.GetAllUser()

	c.JSON(http.StatusOK, user)
}

func GetOneUser(c *gin.Context) {
	user := services.GetOneUser(c.Param("id"))

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	updateUser := entity.User{}

	c.BindJSON(&updateUser)

	user := services.UpdateUser(c.Param("id"), &updateUser)

	c.JSON(http.StatusOK, user)
}
