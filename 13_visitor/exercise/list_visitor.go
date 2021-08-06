package visitor

import "fmt"

type ListVisitor struct {}

func (a *ListVisitor) Visit(e Element) {
	switch e.(type) {
	case *File:
		fmt.Printf("visit file, filename: %s \n", e.(*File).GetName())
	case *Directory:
		d := e.(*Directory)
		fmt.Printf("visit dir, dirname: %s \n", d.GetName())
		for _, child := range d.childs { // todo 这里还可以升级成iterator模式
			child.Accept(a)
		}
	default:
		fmt.Println("err arg")
	}
}
