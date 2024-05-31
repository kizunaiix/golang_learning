package test

import (
	"fmt"
	"golang_learning/pkg/mymath"
	"golang_learning/pkg/objs"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
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
	for k, v := range list1 {
		fmt.Print(k, v+"\n")
	}
	fmt.Println("------------")
	var aa int32 = '算'
	fmt.Println(string(aa)) // 注意这里的string()方法
	fmt.Printf("aa : %v\n", aa)
	fmt.Println("------------")
	s := "Go语言编程"
	//  转成 rune 数组，需要几个字符，取几个字符
	fmt.Println([]rune(s)[:4]) // 输出：Go语言 -> 不正确是因为我没使用string()方法

}

// 数组切片和字符串切片的不同
func Test6(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := "abc"
	slice_a := a[:2]
	slice_b := b[:2]
	slice_aa := slice_a[:1] //切片也能切

	a[1] = 4  //分别改动原数组或字符串
	b = "def" //分别改动原数组或字符串

	fmt.Println(slice_a) //数组的切片也会变
	fmt.Println(slice_b) //字符串的切片不受影响
	fmt.Printf("a=%v,b=%v,slice_aa=%v\n", a, b, slice_aa)
}

func Test7(t *testing.T) {
	type Student struct {
		Age int
	}

	kvmap := map[string]Student{"menglu": {Age: 21}}
	kvmap["menglu"] = Student{Age: 22}
	s := []Student{{Age: 21}}
	s[0].Age = 22
	fmt.Println(kvmap, s)
}

func Test8(t *testing.T) {
	type Student struct {
		Age int
	}

	kvslice := []Student{{Age: 21}}
	fmt.Println(kvslice[0].Age)
}

func Test9(t *testing.T) {
	a := rand.Intn(-1)
	fmt.Println(a)
}

// switch case语句
func Test10(t *testing.T) {
	a := 10
	switch {
	case a == 10:
		println("a is 10")
	case a == 9:
		println("a is 9")
	}
}

// 一个for循环
func Test11(t *testing.T) {
	v := []int{1, 2, 3}
	for k, v := range v {
		println("k:", k, "v:", v)
	}
	println("------------------")
	for i := 0; i < 3; i += 3 {
		println("ok")
	}
}

// 写一个Student的Option模式
func Test12(t *testing.T) {
	stu1 := objs.NewStudent(
		objs.WithStudentAge(23),
		objs.WithStudentId(1145141919810),
		objs.WithStudentName("李天梭"),
	)

	fmt.Printf("学生是：%v", *stu1)
	fmt.Printf("\n其姓名为： %v", stu1.Name)
	fmt.Printf("\n其年龄为： %v + 1 = %v", stu1.Age, stu1.Age+1)
	fmt.Printf("\n其学号为： %v", stu1.Id)
}

// 看一下map是不是指针
func Test13(t *testing.T) {
	aMap := make(map[string]any)
	aMap["age"] = 3
	aMap["name"] = "xiaoming"
	aMap["price"] = 135
	aMap["color"] = "black"
	aMap["color"] = "white"
	bMap := aMap
	bMap["age"] = 8848
	fmt.Println(aMap)
	fmt.Println(bMap)
}

func ModSlice(s *[]string) {
	*s = (*s)[0:2]
}

// 研究一下slice的指针
func Test14(t *testing.T) {

	a := []string{"aa", "bb", "cc", "dd"}
	// a = a[0:2]
	ModSlice(&a)
	fmt.Println(a)
}

/*
试一下接口的方法具体会怎么实现，是当成猫还是狗

	答案是会直接panic
*/
func Test15(t *testing.T) {
	var a objs.Animal
	objs.IsAnimal(a)

	// 用defer把这个panic给recover了, defer后面不是直接写执行语句而是用了匿名函数，这与defer的机制有关：
	// defer后直接写语句，其实会当时就执行这些语句，只是defer之后再显示结果。而defer后用匿名函数就可以保证这些语句一定会等到函数退出前才执行。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	a.Bark()
}

// slice 可以是空，也可以是nil，它们不太一样。
// map 同理。
func Test16(t *testing.T) {
	var nulls = []int32{}
	var nils []int32
	var mymap map[int32]int32

	if len(nulls) == len(nils) && len(mymap) == 0 {
		fmt.Println("空slice和nil slice的长度都是0 , nil的map其长度也是0")
	} else {
		fmt.Println("难道不是吗？")
	}
}

// 对值为nil的map添加键值对，会引发panic
func Test17(t *testing.T) {
	var NilMap map[string]string
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	NilMap["aa"] = "bb"
}

func Test18(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r, " 其类型为： ", reflect.TypeOf(r))
		}
	}()

	panic([]int{1, 2, 34}) // 抛出一个某种类型的异常
}

type Fanxing interface {
	~int | string
}

func UseFanxing[T Fanxing](a T) {
	fmt.Println(a)
}

// 泛型
func Test19(t *testing.T) {
	// UseFanxing(true)    -> 这句会报错，因为泛型里没bool
	UseFanxing(42)
	UseFanxing("ok")
}

// 通道,不使用goroutine
func Test20(t *testing.T) {
	var c = make(chan int, 1) //这里必须容量大于等于2，否则死锁，第255行就无法执行了
	c <- 98
	// c <- 87
	a := <-c
	fmt.Println(a)
}

// 通道，使用goroutine
//
// 主goroutine不可以永久阻塞（会死锁）。别的goroutine可以阻塞，反正最后主函数执行完了会全退出。
func Test21(t *testing.T) {
	var c = make(chan int)

	// 启动一个goroutine来接收数据
	go func() {

		c <- 98
		c <- 87
	}()

	time.Sleep(3 * time.Second)

	a := <-c
	fmt.Println(a)
	time.Sleep(time.Second)
	// 读取第二个值
	b := <-c
	fmt.Println(b)
}

// 把yaml文件转换json文件
func Test22(t *testing.T) {
	err := objs.YamlToJson("../cmd/tt.yml", "../cmd/tt.json")
	if err != nil {
		log.Println(err)
	}
}
