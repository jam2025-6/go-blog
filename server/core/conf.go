package core

import (
	"fmt"
	"log"
	"server/config"
	"server/utils"

	"gopkg.in/yaml.v3"
)

func InitConf() *config.Config {
	c := &config.Config{}
	// 1. 读取 YAML 文件
	yamlConf, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// 2. 解析 YAML 内容
	err = yaml.Unmarshal(yamlConf, c)
	fmt.Println(c)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML configuration: %v", err)
	}
	return c
}
