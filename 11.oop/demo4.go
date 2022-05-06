package main

import (
	"fmt"
)

// interface 实现多态 (本质是指针, 定义接口. 其他类实现相应函数)

type AnimalIF interface {
	Sleep()
	GetColor() string
}

// 定义类
type Cat struct {
	Color string
}

func (this *Cat) Sleep() { fmt.Println("cat is sleep!") }

func (this *Cat) GetColor() string { return this.Color }

// 定义类
type Dog struct {
	Color string
}

func (this *Dog) Sleep() { fmt.Println("dog is sleep!") }

func (this *Dog) GetColor() string { return this.Color }

func DoSomething(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color =", animal.GetColor())
}

func main() {
	var animal AnimalIF
	animal = &Dog{"Green"}
	animal.Sleep()
	animal = &Cat{"Yellow"}
	animal.Sleep()

	dog := Dog{"Green"}
	cat := Cat{"Yellow"}
	DoSomething(&dog)
	DoSomething(&cat)
}
