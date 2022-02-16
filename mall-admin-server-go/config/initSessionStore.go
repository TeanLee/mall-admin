package config

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func InitSessionStore() sessions.Store {
	addr := GetEnv().Db.Redis.Addr
	password := GetEnv().Db.Redis.Password
	sessionSize := GetEnv().Db.Redis.SessionSize
	store, err := redis.NewStore(sessionSize, "tcp", addr, password, []byte("store"))

	if err != nil {
		panic(err)
	}

	// 将 redis 中存储的 session 信息从默认的 gob 改成 json 格式（便于查看，但是传输效率会下降）
	//err, reStore := redis.GetRedisStore(store)
	//if err != nil {
	//	panic(err)
	//}
	//reStore.SetSerializer(redistore.JSONSerializer{})

	return store
}
