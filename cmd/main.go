package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	list1 := [...]string{"abc", "def", "ghi"}
	slice1 := list1[:2]
	fmt.Printf("slice1 : %v \n", slice1)

}
