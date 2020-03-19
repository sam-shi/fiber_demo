package common

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//解析yml文件
type BaseInfo struct {
	MysqlHost     string `yaml:"mysql_host"`
	MysqlPort     string `yaml:"mysql_port"`
	MysqlUsername string `yaml:"mysql_username"`
	MysqlPassword string `yaml:"mysql_password"`
	MysqlDatabase string `yaml:"mysql_database"`
	MysqlChatSet  string `yaml:"mysql_chatset"`
}

func (c *BaseInfo) GetConf() *BaseInfo {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
