package main

import (
	"golang_learning/pkg/objs"
	"log"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world!"))
	// })
	// http.ListenAndServe("localhost:8080", nil)

	err := objs.YamlToJson("../cmd/tt.yml", "../cmd/tt.json")
	if err != nil {
		log.Println(err)
	}

}
