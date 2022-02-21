package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
)

type CategoryAPI struct {
	CategorySrv service.CategorySrv
}

func (CategoryAPI) GetCategories(c *gin.Context) {
	categories, err := service.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
	}
	c.JSON(http.StatusOK, gin.H{"data": categories})
}

func (CategoryAPI) UpdateCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "请说明要修改的 Category Id"})
		return
	}
	var params model.Category
	err = c.Bind(&params)
	if err := service.UpdateCategory(id, params); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
	return
}
