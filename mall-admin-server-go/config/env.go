package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//const configYaml = "/env/local.yaml"

type Mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db-name"`
	Parameters string `yaml:"parameters"`
}

type Redis struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	SessionSize int    `yaml:"session-size"`
}

type Env struct {
	Db struct {
		Mysql `yaml:"mysql"`
		Redis `yaml:"redis"`
	} `yaml:"db"`
}

func (a Mysql) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

func GetEnv() Env {
	t := Env{}

	data, err := os.ReadFile("/Users/litingting/front-end/mall-admin/mall-admin-server-go/env/local.yaml")

	// 根据进程中的环境变量参数，判断使用哪份配置文件
	currentENV := os.Getenv("env")
	fmt.Println("currentENV", currentENV)
	if currentENV == "production" {
		data, err = os.ReadFile(GetAppPath() + "/env/prodEnv.yaml")
	} else if currentENV == "staging" {
		data, err = os.ReadFile(GetAppPath() + "/env/devEnv.yaml")
	}

	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &t)
	if err != nil {
		panic(err)
		return t
	}
	fmt.Println(t.Db.DSN())
	return t
}

// GetAppPath 解决执行 go build 时，获取不到相对路径的问题
func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
