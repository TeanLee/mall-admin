package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/middleware"
	"mall-admin-server-go/router"
)

func main() {
	r := gin.Default()

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(""))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions("mysession", store))

	r.Use(middleware.Cors())
	r2 := router.Router{}
	r2.RegisterAPI(r)
	r.Run()
}
