package controller

import (
	"net/http"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router  *gin.Engine
	service service.AuthService
}

func NewAuthController(router *gin.Engine, service service.AuthService) {
	controller := AuthController{
		router:  router,
		service: service,
	}

	v1 := controller.router.Group("api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controller.createHandler)
		}
	}
}

func (a *AuthController) createHandler(c *gin.Context) {
	var payload model.UserCredential

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}

	token, err := a.service.Login(payload.Username, payload.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"token":  token,
	})
}
