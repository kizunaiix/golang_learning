package main

import (
	"fmt"
	"golang_learning/pkg/objs"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// 读取 YAML 文件
	data, err := os.ReadFile("t.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 将 YAML 内容解析为结构体
	var config objs.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// 输出解析结果
	fmt.Printf("Parsed YAML file into struct:\n%+v\n", config)
}
