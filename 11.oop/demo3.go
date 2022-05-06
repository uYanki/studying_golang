package main

import "fmt"

// 函数名/结构体名/结构体名成员名首字母 大写->文件外可访问, 小写->文件外不可访问

type Human struct {
	Name string
	HP   int
}

func (this Human /*值传递*/) SetName1(name string)   { this.Name = name }
func (this *Human /*指针传递*/) SetName2(name string) { this.Name = name }

type Hero struct {
	Human // 类继承
	Level int
}

func (this *Hero) SetLevel(level int) { this.Level = level }

func main() {

	// human := Human{Name: "uyk", HP: 1}
	human := Human{"uyk", 1}
	fmt.Println(human)
	human.SetName1("name1") // 传递对象副本
	fmt.Println(human)
	human.SetName2("name2") // 传递对象本身
	fmt.Println(human)

	fmt.Println("-------------------------")

	hero1 := Hero{Human{"uyk", 1}, 8}
	fmt.Println(hero1)
	hero1.SetName2("hero1") // 父类方法
	hero1.SetLevel(88)      // 子类方法
	fmt.Println(hero1)

	fmt.Println("-------------------------")

	var hero2 Hero
	hero2.Name = "hero2"
	hero2.HP = 100
	hero2.Level = 8
	fmt.Println(hero2)

}
