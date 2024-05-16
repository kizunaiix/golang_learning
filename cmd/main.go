package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// 定义与 YAML 文件对应的结构体
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int64  `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Dbname   string `yaml:"dbname"`
	} `yaml:"database"`
}

func main() {
	// 读取 YAML 文件
	data, err := os.ReadFile("t.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 将 YAML 内容解析为结构体
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// 输出解析结果
	fmt.Printf("Parsed YAML file into struct:\n%+v\n", config)
}
