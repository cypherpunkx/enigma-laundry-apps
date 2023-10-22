package controller

import (
	"errors"
	"net/http"
	"strconv"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/service"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"enigmacamp.com/enigma-laundry-apps/utils/exception"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type EmployeeController struct {
	router  *gin.Engine
	service service.EmployeeService
}

func (s *EmployeeController) createHandler(c *gin.Context) {
	var employee *model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	employee.ID = common.GenerateUUID()

	validate := validator.New()
	if err := validate.Struct(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := s.service.RegisterNewEmployee(employee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Success Create New Employee",
		Data:    data,
	})
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	username := c.GetString("username")

	name := c.DefaultQuery("name", "")
	phoneNumber := c.DefaultQuery("phoneNumber", "")
	address := c.DefaultQuery("address", "")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := e.service.FindAllEmployee(paginationParam, name, phoneNumber, address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Create New List Employee" + username,
		Data:    data,
		Paging:  paging,
	})
}

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")

	employee, err := e.service.GetEmployeeByID(id)

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
		Message: "Success Get Employee with id = " + id,
		Data:    employee,
	})
}

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee

	id := c.Param("id")

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	employee.ID = id

	validate := validator.New()
	if err := validate.Struct(employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := e.service.UpdateEmployeeByID(&employee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Update Employee",
		Data:    data,
	})

}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := e.service.DeleteEmployeeByID(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Delete Employee By ID" + id,
	})

}

func NewEmployeeController(router *gin.Engine, service service.EmployeeService) {
	ctr := &EmployeeController{
		router:  router,
		service: service,
	}

	v1 := ctr.router.Group("/api/v1")
	{
		employees := v1.Group("/employees")
		{
			employees.GET("/", ctr.listHandler)
			employees.GET("/:id", ctr.getHandler)
			employees.POST("/", ctr.createHandler)
			employees.DELETE("/:id", ctr.deleteHandler)
			employees.PUT("/:id", ctr.updateHandler)
		}
	}

}
