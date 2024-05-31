package main

import (
	"golang_learning/pkg/objs"
	"log"
)

func main() {
	err := objs.YamlToJson("../cmd/tt.yml", "../cmd/tt.json")
	if err != nil {
		log.Println(err)
	}
}
