package utils

import (
	"encoding/json"
	"mall-admin-server-go/model"
	"strconv"
)

func SetRedisRoleByUserId(userId int, role model.Role) error {
	roleM, err := json.Marshal(role)
	// 将 "user1-role" 存入 redis。不设置过期时间，在更新用户权限时，也要更新 redis 中的 role 缓存信息
	if err = rdb.Set("user"+strconv.Itoa(userId)+"-role", string(roleM), 0).Err(); err != nil {
		return err
	}
	return nil
}
