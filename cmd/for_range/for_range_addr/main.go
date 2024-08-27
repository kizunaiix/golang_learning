package main

import "fmt"

func main() {
	a := []string{"nihao", "wohao", "dajiahao"}
	for _, str := range a {
		aa := str
		fmt.Println(&str, " : ", str)
		fmt.Println(&aa, " : ", aa)

	}
}
