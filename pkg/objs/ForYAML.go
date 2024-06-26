package objs

import (
	"encoding/json"
	"fmt"
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

func YamlToJson(srcFile, desFile string) error {
	f, err := os.Open(srcFile)
	defer func() error {
		err = f.Close()
		if err != nil {
			return fmt.Errorf("ERROR in close file: %v", err)
		}
		return nil
	}()
	if err != nil {
		return fmt.Errorf("ERROR in open file: %v", err)
	}

	content := make([]byte, 102400)
	contentInGo := make(map[string]any)

	n, err := f.Read(content)
	if err != nil {
		return fmt.Errorf("ERROR in read content: %v", err)
	} else {
		content = content[:n]
	}

	for _, char := range content {
		flag := false
		if char < 32 && char != 10 && char != 13 && char != 9 { // 允许的控制字符: \n, \r, \t
			flag = true
		}
		if flag {
			return fmt.Errorf("YML文件包含无效的控制字符")
		}
	}

	err = yaml.Unmarshal(content, contentInGo)
	if err != nil {
		return fmt.Errorf("ERROR in unmarshalling yaml: %v", err)
	}

	// fmt.Println(string(content))
	// fmt.Println("====================================================")
	// fmt.Println(contentInGo)

	contentInJson, err := json.Marshal(contentInGo)
	if err != nil {
		return fmt.Errorf("ERROR when marshalling to JSON: %v", err)
	}

	// fmt.Println("======================================================")
	// fmt.Println(string(contentInJson))

	err = os.WriteFile(desFile, contentInJson, 0777)
	if err != nil {
		return fmt.Errorf("ERROR in writting files: %v", err)
	}
	return nil
}
