package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

const configYaml = "/Users/litingting/front-end/mall-admin/mall-admin-server-go/env/mysql.yaml"

type Mysql struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	DBName     string `yaml:"db-name"`
	Parameters string `yaml:"parameters"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
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
	data, err := os.ReadFile(configYaml)

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
