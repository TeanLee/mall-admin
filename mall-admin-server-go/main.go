package main

import (
	"github.com/boj/redistore"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"mall-admin-server-go/middleware"
	"mall-admin-server-go/router"
)

func main() {
	r := gin.Default()
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("1"))
	if err != nil {
		panic(err)
	}

	// 将 redis 中存储的 session 信息从默认的 gob 改成 json 格式（便于查看，但是传输效率会下降）
	err, reStore := redis.GetRedisStore(store)
	if err != nil {
		panic(err)
	}
	reStore.SetSerializer(redistore.JSONSerializer{})

	r.Use(sessions.Sessions("mysession", store))

	r.Use(middleware.Cors())
	r2 := router.Router{}
	r2.RegisterAPI(r)
	r.Run()
}
