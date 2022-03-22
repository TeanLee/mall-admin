package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
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
	page, _ := strconv.Atoi(c.Query("page"))

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if page < 0 {
		page = 0
	}

	fmt.Println("pageSize：", pageSize)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := page * pageSize

	orders, count, err := service.GetOrders(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders, "total": count})
	return
}
