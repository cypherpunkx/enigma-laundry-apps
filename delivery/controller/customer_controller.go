package controller

import (
	"errors"
	"net/http"
	"strconv"

	"enigmacamp.com/enigma-laundry-apps/delivery/middleware"
	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/service"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	router  *gin.Engine
	service service.CustomerService
}

func (s *CustomerController) createHandler(c *gin.Context) {
	var customer *model.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	customer.ID = common.GenerateUUID()

	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := s.service.RegisterNewCustomer(customer)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Success Create New Customer",
		Data:    data,
	})
}

func (e *CustomerController) listHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	phoneNumber := c.DefaultQuery("phoneNumber", "")
	address := c.DefaultQuery("address", "")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := e.service.FindAllCustomer(paginationParam, name, phoneNumber, address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Create New List Customer",
		Data:    data,
		Paging:  paging,
	})
}

func (e *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")

	customer, err := e.service.GetCustomerByID(id)

	if err != nil {
		if errors.Is(err, exception.ErrNotFound) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error() + id,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Get Customer with id = " + id,
		Data:    customer,
	})

}

func (e *CustomerController) updateHandler(c *gin.Context) {
	var customer model.Customer

	id := c.Param("id")

	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	customer.ID = id

	validate := validator.New()
	if err := validate.Struct(customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := e.service.UpdateCustomerByID(&customer)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Update Customer",
		Data:    data,
	})
}

func (e *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := e.service.DeleteCustomerByID(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Delete Customer By ID" + id,
	})

}

func NewCustomerController(router *gin.Engine, service service.CustomerService) {
	ctr := &CustomerController{
		router:  router,
		service: service,
	}

	v1 := ctr.router.Group("/api/v1")
	{
		customers := v1.Group("/customers", middleware.AuthMiddleware())
		{
			customers.GET("/", ctr.listHandler)
			customers.GET("/:id", ctr.getHandler)
			customers.POST("/", ctr.createHandler)
			customers.DELETE("/:id", ctr.deleteHandler)
			customers.PUT("/:id", ctr.updateHandler)
		}
	}

}
