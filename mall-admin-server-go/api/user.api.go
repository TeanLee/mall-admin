package api

import (
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/model"
	"mall-admin-server-go/service"
	"net/http"
	"strconv"
)

type UserAPI struct {
	UserSrv service.UserSrv
}

func (UserAPI) GetUsers(c *gin.Context) {
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

	username := c.Query("username")
	address := c.Query("address")
	phone := c.Query("phone")
	receiver := c.Query("receiver")

	users, count, err := service.GetUsers(offset, pageSize, username, address, phone, receiver)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed get resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": users, "total": count})
	return
}

func (UserAPI) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "请说明要修改的 User Id"})
		return
	}
	var user model.User
	err = c.Bind(&user)
	if err := service.UpdateUser(id, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed update resource", "detail": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
	return
}

func (UserAPI) AddUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "参数有误"})
		return
	}

	_, err := service.AddUser(user)
	if err != nil {
		return
	}
}
