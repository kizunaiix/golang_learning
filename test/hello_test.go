package test

import (
	"fmt"
	"golang_learning/pkg/mymath"
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
	const aa int = 77
	// aa = 55 error
	a, b, c := 5, 7, 4
	print(a + b)
	fmt.Printf("a+b-c-d-f=%d,aa=%d, + %s"+" leiho", (a + b - c - d - f), aa, "haha\n")
}

func Test3(t *testing.T) {
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}
