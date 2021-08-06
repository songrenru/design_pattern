package composite

import "fmt"

//Component ,entry
type Component interface {
	GetName() string
	GetSize() int
	PrintList()
	Add(child Component)
}

const (
	TypeLeaf = iota
	TypeComposite
)

func NewComponent(t int, name string, size int) Component {
	switch t {
	case TypeLeaf:
		l := &Leaf{}
		l.name = name
		l.size = size
		return l
	case TypeComposite:
		c := &Composite{}
		c.name = name
		c.size = size
		return c
	}

	return nil
}

// leaf和composite的parent
type component struct {
	parent Component
	name string
	size int
}

func (c *component) GetName() string {
	return c.name
}

func (c *component) GetSize() int {
	return c.size
}

func (c *component) PrintList() {
	fmt.Printf("%s-%d \n", c.name, c.size)
}

func (c *component) Add(child Component) {
	panic("非法操作")
}

// Leaf
type Leaf struct {
	component
}

// Composite
type Composite struct {
	component
	childs []Component
}

func (c *Composite) PrintList() {
	fmt.Printf("%s-%d \n", c.name, c.size)
	for _, child := range c.childs {
		child.PrintList()
	}
}

func (c *Composite) Add(child Component) {
	c.childs = append(c.childs, child)
	c.size += child.GetSize()
	// todo 还是要一个父component的struct，来统计name、size，既方便，也能实现child添加时size的累加
}


