package objs

import "fmt"

type Animal interface {
	Bark()
}

type Dog struct {
	Name string
}

func (d Dog) bark() {
	fmt.Println("wang wang wang")
}

type Cat struct {
	Name string
}

func (c Cat) Bark() {
	fmt.Println("miao miao")
}

func IsAnimal(a Animal) {
	fmt.Println("is an animal!")
}
