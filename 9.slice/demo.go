package main

import "fmt"

func printArray(arr []int /*指针传递*/) {
	// go语言不允许存在不使用的变量, 若不使用需用'_'代替
	for _, value := range arr {
		fmt.Println("value =", value)
	}
	arr[0] = 666 // 修改元素值
}

// 动态数组 slice (切片)
func main() {
	array := []int{1, 2, 3, 4}
	fmt.Printf("type of array = %T\n", array)
	printArray(array)
	fmt.Println("array[0] =", array[0]) // 值被改变
}
