package visitor

import "fmt"

// Visitor
type Visitor interface {
	Visit(e Element)
}

// acceptor
type Element interface {
	Accept(v Visitor)
}

// todo 虽然抽象了Visitor，acceptor，但是具体visitor需要识别element的类型，进行不同的visit逻辑（这样是否丧失了解耦的意义呢？）

type VisitorA struct {}

func (a *VisitorA) Visit(e Element) {
	switch e.(type) {
	case *ElementA:
		fmt.Printf("A visit ele: %s, 并给了他两拳 \n", e.(*ElementA).Name)
	default:
		fmt.Println("err arg")
	}
}

type VisitorB struct {}

func (b *VisitorB) Visit(e Element) {
	switch e.(type) {
	case *ElementA:
		fmt.Printf("B visit ele: %s, 并亲了他两口 \n", e.(*ElementA).Name)
	default:
		fmt.Println("err arg")
	}
}

type ElementA struct {
	Name string
}

func (a *ElementA) Accept(v Visitor) {
	v.Visit(a)
}


