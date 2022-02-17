package utils

import (
	"encoding/json"
	"mall-admin-server-go/config"
	"mall-admin-server-go/model"
	"strconv"
)

var rdb = config.InitRedis()

func GetUserRoleFromRedisByUserId(userId int) (model.Role, error) {
	var (
		role        model.Role
		roleFromRdb string
		err         error
	)
	if roleFromRdb, err = rdb.Get("user" + strconv.Itoa(userId) + "-role").Result(); err != nil {
		return role, err
	}
	if err := json.Unmarshal([]byte(roleFromRdb), &role); err != nil {
		return role, err
	}
	return role, nil
}
