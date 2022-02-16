package config

import (
	"github.com/go-redis/redis"
)

// InitRedis 暂时没有用到。对 session 的 操作，有 store 的方法
func InitRedis() *redis.Client {
	addr := GetEnv().Db.Redis.Addr
	password := GetEnv().Db.Redis.Password

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		return nil
	}
	return rdb
}

/**

//存值(key, value, expire)
if err := dao.Redis.Set(key, val, 0).Err(); err != nil{
fmt.Println(err)
}

//取值(key, value, expire)
if err := dao.Redis.Get(key).Err(); err != nil{
fmt.Println(err)

*/
