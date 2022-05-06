package main

import "fmt"

// 类似枚举类型
const (
	a = iota * 2 // 行索引'iota'(从0开始)
	b
	c
	d = iota * 3
	e
	f
)

const (
	aa, bb = iota * 2, iota * 3
	cc, dd
	ee, ff = iota - 1, iota - 2
	gg, hh
)

func main() {
	const length int = 10 // 常量
	fmt.Println("length =", length)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)

	fmt.Println("d =", d)
	fmt.Println("e =", e)
	fmt.Println("f =", f)

	fmt.Println("aa =", aa)
	fmt.Println("bb =", bb)
	fmt.Println("cc =", cc)
	fmt.Println("dd =", dd)

	fmt.Println("ee =", ee)
	fmt.Println("ff =", ff)
	fmt.Println("gg =", gg)
	fmt.Println("hh =", hh)
}
