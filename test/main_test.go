package test

import (
	"fmt"
	"golang_learning/pkg/mymath"
	"golang_learning/pkg/objs"
	"log"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
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

// 试一下接口的方法具体会怎么实现，是当成猫还是狗
//
//	答案是会直接panic
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

// Worker函数，负责从任务通道接收任务并处理
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // 模拟工作耗时
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // 处理结果
	}
}

// 学习goroutine ，这里用的通道有缓冲。
func Test16(t *testing.T) {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup

	// 启动worker goroutine
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 发送任务到任务通道
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 所有任务发送完毕后关闭通道

	// 等待所有worker完成
	wg.Wait()
	close(results)

	// 收集处理结果
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// 试一下0缓冲的通道
func Test17(t *testing.T) {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int)    // 无缓冲通道
	results := make(chan int) // 无缓冲通道
	var wg sync.WaitGroup

	// 启动worker goroutine
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// 向jobs通道发送任务
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// 等待所有worker完成任务
	// 这里要用一个goroutine来执行wg.Wait()，这是因为主goroutine上一直堵着for range
	// 因为主goroutine堵住了，所以要让一个新的goroutine来执行wg.Wait()；close(results)
	go func() {
		wg.Wait()
		close(results)
	}()

	// 从results通道接收并打印结果，
	// 这里的for range对通道作用的效果是，会一直阻塞，直到通道关闭。
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}

// YAML文件读取和unmarshal
func Test18(t *testing.T) {
	// 读取 YAML 文件
	data, err := os.ReadFile("t.yml")
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// 将 YAML 内容解析为结构体
	var config objs.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	// 输出解析结果
	fmt.Printf("Parsed YAML file into struct:\n%+v\n", config)
}

type weirdPeople[T killer] struct{} // 注意泛型的写法， 定义类型的时候就得写上泛型了

func (weirdPeople[T]) weird(t T) { // 注意泛型的写法， 绑定方法的时候直接使用泛型
	fmt.Println("werid: ", t)
}

type killer interface {
	int | string
	// kill()        ----> 这里不能写kill方法  ，因为写kill的话，int和string就必须得会kill。 当然如果不是int和string而是people这种类型的话那没问题。
}

// 试试interface什么时候必须写作泛型
func Test19(t *testing.T) {
	var m int = 3
	var n weirdPeople[int] //  声明变量的时候必须实例化泛型
	n.weird(m)             // m作为int符合weird()的要求。

}
