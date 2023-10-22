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

type ProductController struct {
	router  *gin.Engine
	service service.ProductService
}

func (s *ProductController) createHandler(c *gin.Context) {
	var product *model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	product.ID = common.GenerateUUID()

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := s.service.RegisterNewProduct(product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Success Create New Product",
		Data:    data,
	})
}

func (e *ProductController) listHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := e.service.FindAllProduct(paginationParam, name)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Create New List Product",
		Data:    data,
		Paging:  paging,
	})
}

func (e *ProductController) getHandler(c *gin.Context) {
	id := c.Param("id")

	product, err := e.service.GetProductByID(id)

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

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Get Product with id = " + id,
		Data:    product,
	})

}

func (e *ProductController) deleteHandler(c *gin.Context) {
	id := c.Param("id")

	if err := e.service.DeleteProductByID(id); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Delete Product By ID" + id,
	})

}

func (e *ProductController) updateHandler(c *gin.Context) {
	var product model.Product

	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	product.ID = id

	validate := validator.New()
	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	data, err := e.service.UpdateProductByID(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusOK,
		Message: "Success Update Product",
		Data:    data,
	})

}

func NewProductController(router *gin.Engine, service service.ProductService) {
	ctr := &ProductController{
		router:  router,
		service: service,
	}

	v1 := ctr.router.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("/", middleware.AuthMiddleware(), ctr.listHandler)
			products.GET("/:id", ctr.getHandler)
			products.POST("/", ctr.createHandler)
			products.DELETE("/:id", ctr.deleteHandler)
			products.PUT("/:id", ctr.updateHandler)
		}
	}

}
