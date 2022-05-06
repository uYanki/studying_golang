package main

import "fmt"

func foo1() { fmt.Println("foo1") }
func foo2() { fmt.Println("foo2") }

func main() {

	// 在函数体结束之前触发 defer (类似c++的析构函数)
	// 多 defer 的调用顺序, 类似于栈, 后压栈的先出栈, 即后定义的先调用
	defer fmt.Println("bye Go!")
	defer fmt.Println("bye*2 Go!")
	defer foo1()
	defer foo2()

	fmt.Println("hello Go!")

}
