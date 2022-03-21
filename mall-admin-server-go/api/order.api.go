package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/service"
	"net/http"
)

type OrderAPI struct {
	OrderSrv service.OrderSrv
}

type QueryOrderParam struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"`
	ProductName string `json:"productName"`
	Category    string `json:"category"`
}

func (OrderAPI) GetOrders(c *gin.Context) {
	var param QueryOrderParam
	err := c.BindJSON(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误", "detail": err})
		return
	}
	page := param.Page
	fmt.Println("page：", page)
	if page < 0 {
		page = 0
	}

	pageSize := param.PageSize
	fmt.Println("pageSize：", pageSize)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := page * pageSize

	productName := param.ProductName
	category := param.Category

	orders, count, err := service.GetOrders(offset, pageSize, productName, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders, "total": count})
	return
}
