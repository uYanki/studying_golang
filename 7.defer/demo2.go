package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}
func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func ReturnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	// 运行结果: defer 在 return 之后调用
	ReturnAndDefer()
}
