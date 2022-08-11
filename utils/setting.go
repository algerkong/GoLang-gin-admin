package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败, 请检查!", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
    JwtKey = file.Section("server").Key("HttpPort").MustString("algerkong")
}

func LoadData(file *ini.File) {
	Db = file.Section("data").Key("Db").MustString("mysql")
	DbHost = file.Section("data").Key("DbHost").MustString("localhost")
	DbPort = file.Section("data").Key("DbPort").MustString("3306")
	DbUser = file.Section("data").Key("DbUser").MustString("root")
	DbPassword = file.Section("data").Key("DbPassword").MustString("alger")
	DbName = file.Section("data").Key("DbName").MustString("ginblog")
}
