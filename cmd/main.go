package main

import (
	// import的写法是  ModuleName/DirName/DirName  导入的是package  用的时候写PackageName.Func()
	"golang_learning/pkg/objs"
	"log"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("hello world!"))
	// })
	// http.ListenAndServe("localhost:8080", nil)

	err := objs.YamlToJson("../cmd/tt1.yml", "../cmd/tt.json")
	if err != nil {
		log.Println(err)
	}

}
