package main

import "fmt"

// 结构体
type Book struct {
	title  string
	author string
}

func changeBookName1(book Book)  { book.title = "111" } // 值传递
func changeBookName2(book *Book) { (*book).title = "222" }
func changeBookName3(book *Book) { book.title = "333" } // 可以省略解引用'*'

func main() {
	var book Book
	book.title = "title"
	book.author = "author"
	fmt.Printf("%v\n", book)

	changeBookName1(book)
	fmt.Printf("%v\n", book)

	changeBookName2(&book)
	fmt.Printf("%v\n", book)

	changeBookName3(&book)
	fmt.Printf("%v\n", book)
}
