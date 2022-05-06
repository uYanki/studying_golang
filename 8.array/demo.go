package main

import "fmt"

func printArray(arr [4]int /*值传递*/) {
	for i := 0; i < len(arr); i++ {
		fmt.Println("value =", arr[i])
	}
	arr[0] = 666
}

// 静态数组
func main() {
	var array1 [8]int                  // 元素值默认为0
	for i := 0; i < len(array1); i++ { // 索引从0开始, 且没有 ++i
		fmt.Println("value =", array1[i])
	}
	fmt.Println("---------")
	array2 := [8]int{1, 2, 3, 4}       // 前4个元素值为提供的值, 剩余的值为0
	for index, value := range array2 { // 类似 c++ 的 for_each, python 的 enumerate
		fmt.Println(index, value)
	}
	fmt.Println("---------")
	array3 := [4]int{1, 2, 3, 4}
	printArray(array3)
	fmt.Println("array[0] =", array3[0]) // 值传递不改变值
	fmt.Println("---------")
	fmt.Printf("type of array1 = %T\n", array1)
	fmt.Printf("type of array2 = %T\n", array2)
	fmt.Printf("type of array3 = %T\n", array3)
	// printArray(array1) // 与 c/c++(传递指针) 不同, [10]int 不能作为 [4]int 的参数
}
