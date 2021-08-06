package main

import (
	"fmt"
	iterator "github.com/songrenru/design_attern/12_iterator"
)

func main() {
	bookshelf := iterator.NewBookshelf()
	bookshelf.AddBook(iterator.NewBook("book1"))
	bookshelf.AddBook(iterator.NewBook("book2"))
	bookshelf.AddBook(iterator.NewBook("book3"))
	bookshelf.AddBook(iterator.NewBook("book4"))

	iterator1 := bookshelf.Iterator()
	for iterator1.HasNext() {
		book := iterator1.Next().(*iterator.Book)
		fmt.Println(book)
	}
}
