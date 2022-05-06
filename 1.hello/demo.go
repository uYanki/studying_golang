package main

import (
	"fmt"
	"time"
)

// ① 函数名与'{'必须同行
// ② 语句后的';',可加可不加
func main() {
	fmt.Println("hello Go!")
	time.Sleep(2 * time.Second)
	fmt.Println("hello Go!")
}
