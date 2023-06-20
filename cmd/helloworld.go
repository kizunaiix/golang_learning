package main

import (
	"fmt"
	"golang_learning/pkg/mymath"
)

func main() {

	fmt.Println("hello world")
	s := mymath.Add(3, 4)
	fmt.Printf("s = %v\n", s)
	fmt.Println("------")
	fmt.Println(intAdd7(2, 3))
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}

func intAdd7(a int, b int) int {
	var c int = 7
	var gg float64 = 2.3
	f := func(o int) int { return o }
	return f(c) + b + int(gg)
}
