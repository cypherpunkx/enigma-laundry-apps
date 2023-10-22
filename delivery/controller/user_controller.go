package controller

import (
	"net/http"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/service"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router  *gin.Engine
	service service.UserService
}

func NewUserController(router *gin.Engine, service service.UserService) {
	controller := UserController{
		router:  router,
		service: service,
	}

	v1 := controller.router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/", controller.createHandler)
		}
	}
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user.ID = common.GenerateUUID()

	if err := u.service.RegisterNewUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "Successfully Create User",
	})
}
