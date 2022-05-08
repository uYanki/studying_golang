package main

import "fmt"

// 万能类型 interface{}
func foo(arg interface{}) {

	// pair<type,value>
	fmt.Println(arg)
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg'type isn't string")
	} else {
		fmt.Println("arg'type is string:", value)
	}
}

func main() {
	foo(0)
	foo("ok")
}
