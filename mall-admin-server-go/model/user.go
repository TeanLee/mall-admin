package model

type User struct {
	Uid      int    `json:"uid"`
	Username int    `json:"username"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Receiver string `json:"receiver"`
}

func (User) TableName() string {
	return "user"
}
