package main

import "fmt"

func main() {
	print("hello world")
	// var c rune = 65
	xx := fmt.Sprintf("a"+"%v"+"b", 65)
	fmt.Println(xx)

}
