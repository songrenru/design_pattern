package visitor

import "fmt"

type Entry interface {
	Element
	GetName() string
	GetSize() int
	PrintList()
	Add(child Entry)
}

const (
	TypeLeaf = iota
	TypeComposite
)

func NewEntry(t int, name string, size int) Entry {
	switch t {
	case TypeLeaf:
		l := &File{}
		l.name = name
		l.size = size
		return l
	case TypeComposite:
		c := &Directory{}
		c.name = name
		c.size = size
		return c
	}

	return nil
}

// Leaf
type File struct {
	parent Entry
	name string
	size int
}

func (c *File) GetName() string {
	return c.name
}

func (c *File) GetSize() int {
	return c.size
}

func (c *File) PrintList() {
	fmt.Printf("%s-%d \n", c.name, c.size)
}

func (c *File) Add(child Entry) {
	panic("非法操作")
}

func (a *File) Accept(v Visitor) {
	v.Visit(a)
}

// Component
type Directory struct {
	parent Entry
	name string
	size int
	childs []Entry
}

func (c *Directory) GetName() string {
	return c.name
}

func (c *Directory) GetSize() int {
	return c.size
}

func (c *Directory) PrintList() {
	fmt.Printf("%s-%d \n", c.name, c.size)
	for _, child := range c.childs {
		child.PrintList()
	}
}

func (c *Directory) Add(child Entry) {
	c.childs = append(c.childs, child)
	c.size += child.GetSize()
}

func (a *Directory) Accept(v Visitor) {
	v.Visit(a)
}



