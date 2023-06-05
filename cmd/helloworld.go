package main

import (
	"fmt"
	"golang_learning/pkg/mymath"
)

func main() {
	fmt.Println("hello world")
	s := mymath.Add(3, 4)
	fmt.Printf("s = %v", s)
}
