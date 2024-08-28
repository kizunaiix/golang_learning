package main

import "fmt"

func main() {
	a := []string{"nihao", "wohao", "dajiahao"}
	for _, str := range a {
		aa := str // 从1.22版本之后就不用再加这一句了
		fmt.Println("str-> ", &str, " : ", str)
		fmt.Println("aa-> ", &aa, " : ", aa)

	}
}
