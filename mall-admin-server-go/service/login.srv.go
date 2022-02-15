package service

type LoginSrv struct {
	userName string
	password string
	role     string
}

type LoginBody struct {
	userName string `json:"username"`
	password string `json:"password"`
}

func (a LoginSrv) Verify(username string, password string) bool {
	if username == "admin" && password == "admin" {
		return true
	} else {
		return false
	}
}

func Verify(username string, password string) bool {
	if username == "admin" && password == "admin" {
		return true
	} else {
		return false
	}
}
