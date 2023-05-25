package main

import (
	"fmt"
	"golang_learning/mymath"
	"testing"
)

var d int = 9
var f = 10

// 试一下import与直接print
func Test1(t *testing.T) {
	s := 3
	println(s)
	print(mymath.Add(3, 4))
}

// 没啥
// @param:
func Test2(t *testing.T) {
	var a int = 3
	a, b, c := 5, 7, 4
	print(a + b)
	fmt.Printf("%d, + %s", (a + b - c), "haha\n")
}
