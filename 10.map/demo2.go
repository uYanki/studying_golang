package main

import "fmt"

func changeValue(mymap map[string]string /*指针传递*/) { mymap["one"] = "null" }

func main() {
	mymap := make(map[string]string)

	// 添加
	mymap["one"] = "java"
	mymap["two"] = "c++"
	mymap["three"] = "pyhton"

	// 遍历
	for key, value := range mymap {
		fmt.Println(key, value)
	}

	// 删除
	delete(mymap, "three")

	// 修改
	mymap["two"] = "php"
	changeValue(mymap)

	fmt.Println("-----------------------------------")

	for key, value := range mymap {
		fmt.Println(key, value)
	}
}
