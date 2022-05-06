package main

import "fmt"

type myint int // 类型别名 类似 c/c++ 的 typedef

func main() {
	var num myint = 1
	fmt.Printf("%T\n", num)
}
