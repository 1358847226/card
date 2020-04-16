package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	// 站点配置
	App struct {
		Name  string
		Port  string
		Debug bool
		Url   string
		Tcp  string
	}
	//数据库连接配置
	Conenction struct {
		Type     string
		Host     string
		Port	 string
		DataBase string
		Username string
		Password string
	}
}

var AppConfig Config

func init() {
	file, err := ioutil.ReadFile("./config/conf.yaml")
	if err != nil {
		log.Panicln("加载配置文件失败", err)
	}

	if err := yaml.Unmarshal(file, &AppConfig); err != nil {
		log.Fatalf("解析配置文件失败:", err)
	}
}