package main

import (
	"fmt"
	"golang_learning/pkg/objs"
)

func main() {

	// fmt.Println("hello world")
	// s := mymath.Add(3, 4)
	// fmt.Printf("s = %v\n", s)
	// fmt.Println("------")
	// fmt.Println(intAdd7(2, 3))
	// fmt.Println("Press Enter to exit...")
	// fmt.Scanln()
	// // leecode.Merge()
	stu1 := objs.NewStudent(
		objs.WithStudentAge(23),
		objs.WithStudentId(1145141919810),
		objs.WithStudentName("李天梭"),
	)

	fmt.Println(*stu1)
}

func intAdd7(a int, b int) int {
	var c int = 7
	var gg float64 = 2.3
	f := func(o int) int { return o }
	return f(c) + a + b + int(gg)
}
