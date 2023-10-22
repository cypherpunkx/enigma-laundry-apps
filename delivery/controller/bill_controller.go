package controller

import (
	"net/http"
	"strconv"

	"enigmacamp.com/enigma-laundry-apps/model"
	"enigmacamp.com/enigma-laundry-apps/model/dto"
	"enigmacamp.com/enigma-laundry-apps/service"
	"enigmacamp.com/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type BillController struct {
	router      *gin.Engine
	billService service.BillService
}

func (b *BillController) createHandler(c *gin.Context) {
	bill := model.Bill{}

	if err := c.ShouldBindJSON(&bill); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	bill.ID = common.GenerateUUID()

	if err := b.billService.RegisterNewBill(&bill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	billDetails := []model.BillDetails{}

	for _, item := range bill.BillDetails {
		billDetails = append(billDetails, model.BillDetails{
			ID:         item.ID,
			BillID:     item.BillID,
			ProductID:  item.ProductID,
			FinishDate: item.FinishDate,
		})
	}

	c.JSON(http.StatusCreated, dto.Response{
		Status:  http.StatusCreated,
		Message: "Successfully create a bill",
		Data: model.Bill{
			ID:          bill.ID,
			BillDate:    bill.BillDate,
			EntryDate:   bill.EntryDate,
			CustomerID:  bill.CustomerID,
			EmployeeID:  bill.EmployeeID,
			BillDetails: billDetails,
		},
	})
}

func (b *BillController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := b.billService.ListTransactions(paginationParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Status:  http.StatusOK,
		Message: "Successfully Get All Bills",
		Data:    data,
		Paging:  paging,
	})
}

func (b *BillController) getHandler(c *gin.Context) {
	billID := c.Param("id")

	data, err := b.billService.FindBillByID(billID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Bill",
		"data":    data,
	})
}

func NewBillController(router *gin.Engine, service service.BillService) {
	controller := BillController{
		router:      router,
		billService: service,
	}

	v1 := controller.router.Group("/api/v1")
	{
		bills := v1.Group("/bills")
		{
			bills.POST("/", controller.createHandler)
			bills.GET("/", controller.listHandler)
			bills.GET("/:id", controller.getHandler)
		}
	}
}
