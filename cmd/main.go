package main

import "fmt"

// Barkable 接口
// type Barkable interface {
// 	(any) bark()
// }

type Animal struct {
	Dog
	Duck
}

// Dog 结构体
type Dog struct{}

// Duck 结构体
type Duck struct{}

// Dog 实现 Barkable 接口的 bark 方法
func (d Dog) bark() {
	fmt.Println("wang!")
}

// Duck 实现 Barkable 接口的 bark 方法
// func (d Animal) bark() {
// 	fmt.Println("gua!")
// }

func main() {
	b1 := Animal{}
	b1.bark() // 输出: wang!
}
