package main

import "fmt"

// map 键值对
func main() {

	var map1 map[string]string // map[键类型]值类型
	if map1 == nil {
		fmt.Println("map is nil")
	}
	map1 = make(map[string]string, 2) // 使用前需分配空间
	map1["one"] = "java"
	map1["two"] = "c++"
	map1["three"] = "pyhton"
	fmt.Println(map1) // 输出顺序是按哈希来排序

	map2 := make(map[int]string)
	map2[1] = "java"
	map2[2] = "c++"
	map2[3] = "pyhton"
	fmt.Println(map2)

	map3 := map[string]string{
		"one":   "java",
		"two":   "c++",
		"three": "pyhton",
	}
	fmt.Println(map3)

}
