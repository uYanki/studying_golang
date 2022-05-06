package main

import "fmt"

// 全局变量不能使用':='来定义
var g_a int = 200
var g_b = 200

func main() {

	var a int // 变量默认值为0
	fmt.Println("a =", a)
	fmt.Printf("typen of a = %T\n", a)

	var b int = 100 // 变量默认值
	fmt.Println("b =", b)
	fmt.Printf("typen of b = %T\n", b)

	var c = 100 // 省去数据类型
	fmt.Println("c =", c)
	fmt.Printf("typen of c = %T\n", c)

	d := 100 // 使用':='定义并初始化变量
	fmt.Println("d =", d)
	fmt.Printf("typen of d = %T\n", d)

	// 全局变量
	fmt.Println("g_a =", g_a)
	fmt.Println("g_b =", g_b)

	// 多变量定义
	var e, ee int = 100, 200
	fmt.Println("e =", e, ",", "ee =", ee)
	var f, ff = 100, "str"
	fmt.Println("f =", f, ",", "ff =", ff)

	var (
		g   int     = 100
		gg  bool    = true
		ggg float32 = 3.14
	)
	fmt.Println("g =", g, ",", "gg =", gg, ",", "ggg = ", ggg)

}
