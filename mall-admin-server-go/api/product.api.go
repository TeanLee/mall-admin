package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/config"
	"mall-admin-server-go/model"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
	"strings"
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

	productName := c.Query("productName")
	category := c.Query("category")

	products, count, err := service.GetProducts(offset, pageSize, productName, category)
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

func (ProductAPI) AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数有误"})
		return
	}

	_, err := service.AddProduct(product)
	if err != nil {
		return
	}
}

func (ProductAPI) MostAdd(c *gin.Context) {
	var rdb = config.InitRedis()
	// ZRevRange 取出加入购物车数量前十的商品（分数从高到低）
	productsSet, err := rdb.ZRevRange("products", 0, 10).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err})
		return
	}

	// productsTop map 用于存储 { 商品名：加购数量 }
	productsTop := make(map[string]float64)

	for _, val := range productsSet {
		// 在小程序埋点时，redis 中记录的方式是 product-ID 因此通过 split 方式取出商品 id
		productIdStr := strings.Split(val, "-")[1]
		productId, err := strconv.Atoi(productIdStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err})
			return
		}
		// 通过 productId 找商品名
		product, err := model.GetProductById(productId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err})
			return
		}

		// ZScore 通过 product-id 反找 score （被加购的数量）
		productsTop[product.Title] = rdb.ZScore("products", val).Val()
	}
	c.JSON(http.StatusOK, gin.H{"data": productsTop})

}
