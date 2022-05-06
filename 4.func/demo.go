package main

import "fmt"

func calc_1(a int, b int) (int, int /*匿名返回值*/) { return a + b, a - b }

func calc_2(a int, b int) (add int, sub int) { // 返回值也是形参,其默认值为0
	fmt.Println("add =", add, ",", "sub =", sub)
	add = a + b
	sub = a - b
	fmt.Println("add =", add, ",", "sub =", sub)
	return
}

func main() {

	r1, r2 := calc_1(1, 2)
	fmt.Println("r1 =", r1, ",", "r2 =", r2)

	r3, r4 := calc_2(1, 2)
	fmt.Println("r3 =", r3, ",", "r4 =", r4)

}
