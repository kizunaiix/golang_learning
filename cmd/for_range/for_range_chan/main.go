package main

import (
	"fmt"
	"time"
)

func main() {

	mychan := make(chan string, 3)

	go func() {
		mychan <- "nihaoya"
		mychan <- "woshi"
		mychan <- "kasha"
		mychan <- "nanine?"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		close(mychan)
	}()

	// fmt.Println(<-mychan)
	// fmt.Println(<-mychan)

	// fmt.Println("-----------")

	for a := range mychan {
		fmt.Println(a)
	}
}

// TODO 1. for range chan  2. func(arg)(a,b){}
