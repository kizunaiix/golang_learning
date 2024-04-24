package main

import (
	"fmt"
	"testing"
)

func Test88(t *testing.T) {
	myslice := []int{1, 2, 3, 0, 0, 0}
	merge(myslice, 3, []int{2, 5, 6}, 3)
	fmt.Printf("output:%v", myslice)
}

func Test27(t *testing.T) {
	myslice := []int{3, 2, 2, 3}
	a := removeElement(myslice, 3)
	fmt.Printf("output: %v,%v", a, myslice)
}
