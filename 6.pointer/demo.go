package main

import "fmt"

func swap(pa *int, pb *int) {
	var t int
	t = *pa
	*pa = *pb
	*pb = t
}

func main() {
	var a, b int = 1, 2
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	swap(&a, &b)
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	var p *int = &a   // 一级指针
	var pp **int = &p // 二级指针

	fmt.Println("&a =", &a)
	fmt.Println("p(&a) =", p)
	fmt.Println("pp(&p) =", pp)

}
