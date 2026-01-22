package main

import "fmt"

//定义一个结构体
type T struct {
	name string
}

func (t T) method1() {
	t.name = "new name1"
}

func (t *T) method2() {
	t.name = "new name2"
}

func main() {

	t := T{"old name"}

	fmt.Println("method1 调用前 ", t.name)
	t.method1()
	fmt.Println("method1 调用后 ", t.name)

	fmt.Println("method2 调用前 ", t.name)
	t.method2()
	fmt.Println("method2 调用后 ", t.name)
}

//运行结果：
//method1 调用前  old name
//method1 调用后  old name
//method2 调用前  old name
//method2 调用后  new name2
//从上面的结果可以看出，method1 是值接收者，对结构体的修改不会影响到原始结构体；
// 而 method2 是指针接收者，对结构体的修改会影响到原始结构体。
