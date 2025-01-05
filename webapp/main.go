package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

// Config 结构体定义
type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func loadConfig(filename string) (Config, error) {
	var config Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(file, &config)
	return config, err
}

func main() {
	config, err := loadConfig("config.yaml")
	if err != nil {
		log.Fatalf("无法读取配置文件: %v", err)
	}

	fmt.Printf("Loaded configuration: %+v\n", config)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + config.Server.Port)
}
