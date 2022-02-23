package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
)

type ProductAPI struct {
	ProductSrv service.ProductSrv
}

func (ProductAPI) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 0 {
		page = 0
	}

	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := page * pageSize

	products, count, err := service.GetProducts(offset, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": products, "total": count})
	return
}

func (ProductAPI) UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "请说明要修改的 Product Id"})
		return
	}
	var params model.Product
	err = c.Bind(&params)
	if err := service.UpdateProduct(id, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
	return
}
