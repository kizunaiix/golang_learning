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
	fmt.Printf("s:  %v", s)
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

func helloer() {
	fmt.Println("hello!")
}

// 函数作为变量
func Test4(t *testing.T) {
	s := helloer
	s()
}

func Test5(t *testing.T) {
	list1 := [3]string{"helo"}
	list1[1] = "3"
	fmt.Printf("%v \n", list1)
	fmt.Println(list1)
	fmt.Println("------------")
	var aa int32 = '算'
	fmt.Println(string(aa)) // 注意这里的string()方法
	fmt.Printf("aa : %v\n", aa)
	fmt.Println("------------")
	s := "Go语言编程"
	//  转成 rune 数组，需要几个字符，取几个字符
	fmt.Println([]rune(s)[:4]) // 输出：Go语言 -> 不正确是因为我没使用string()方法

}
func Test6(t *testing.T) {

}
