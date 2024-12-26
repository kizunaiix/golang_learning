package main

import (
	"fmt"
	"sync"
	"time"
)

func MyAsyncFunc(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Println("wake!")
}
func ttt() *string {
	var ss string = "kop"
	return &ss
}
func main() {
	// var wg = &sync.WaitGroup{}
	// fmt.Println("start")
	// wg.Add(1)

	// go MyAsyncFunc(wg)
	// wg.Wait()
	// fmt.Println("end")
	a := ttt()
	fmt.Println(*a)
}
