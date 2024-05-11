package test

import (
	"fmt"
	"golang_learning/leecode"
	"testing"
)

func TestLeecode88(t *testing.T) {
	myslice := []int{1, 2, 3, 0, 0, 0}
	leecode.Merge(myslice, 3, []int{2, 5, 6}, 3)
	fmt.Printf("output:%v", myslice)
}

func TestLeecode27(t *testing.T) {
	myslice := []int{3, 2, 2, 3}
	a := leecode.RemoveElement(myslice, 3)
	fmt.Printf("output: %v,%v", a, myslice)
}
