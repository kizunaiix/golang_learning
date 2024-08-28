package main

import (
	"fmt"
)

// // 定义一个使用泛型的匿名函数  --> 匿名函数不能用泛型
//
//	var add = func [T int | float64](a, b T) T {
//		return a + b
//	}
func add[T int | float64](a, b T) T {
	return a + b
}

func main() {

	// 使用匿名函数
	result1 := add(3, 4)     // 使用int类型
	result2 := add(3.5, 4.5) // 使用float64类型

	fmt.Println(result1) // 输出: 7
	fmt.Println(result2) // 输出: 8.0
}
