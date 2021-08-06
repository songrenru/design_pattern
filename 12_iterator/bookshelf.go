package iterator

import (
	"fmt"
)

// Book
type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}

func (b *Book) GetName() string {
	return b.name
}

func (b *Book) String() string {
	return fmt.Sprintf("book name: %s", b.name)
}

// concrete aggregate
type Bookshelf struct {
	books []*Book
}

func NewBookshelf() *Bookshelf {
	return &Bookshelf{}
}

func (s *Bookshelf) GetBookAt(i int) *Book {
	return s.books[i]
}

func (s *Bookshelf) AddBook(book *Book)  {
	s.books = append(s.books, book)
}

func (s *Bookshelf) GetLenth() int {
	return len(s.books)
}

func (s *Bookshelf) Iterator() Iterator {
	return &bookIterator{bookshelf: s}
}

// concrete iterator
type bookIterator struct {
	bookshelf *Bookshelf
	index int
}

func (i *bookIterator) HasNext() bool {
	return i.index < i.bookshelf.GetLenth()
}

func (i *bookIterator) Next() interface{} {
	res := i.bookshelf.GetBookAt(i.index)
	i.index++
	return res
}






