package service

import (
	"encoding/json"
	"mall-admin-server-go/model"
)

type UserSrv struct{}

func GetUsers(offset int, pageSize int, username string, address string, phone string, receiver string) ([]model.User, int64, error) {
	users, err := model.GetUsersByPagination(offset, pageSize, username, address, phone, receiver)
	total := model.CountUser()
	if err != nil {
		return nil, total, err
	}
	return users, total, nil
}

func GetUserById(userId int) (model.User, error) {
	return model.GetUserById(userId)
}

func UpdateUser(userId int, data interface{}) error {
	marshalData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var user model.User
	err = json.Unmarshal(marshalData, &user)
	if err != nil {
		return err
	}
	user.Uid = userId
	return model.UpdateUser(userId, user)
}

func AddUser(user model.User) (int, error) {
	_, err := model.AddUser(user)
	if err != nil {
		return -1, err
	}
	return 0, nil
}
