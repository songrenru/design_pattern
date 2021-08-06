package main

import visitor "github.com/songrenru/design_attern/13_visitor"

func main() {
	eA := &visitor.ElementA{Name: "eleA"}
	vA := &visitor.VisitorA{}
	vB := &visitor.VisitorB{}

	eA.Accept(vA)
	eA.Accept(vB)
}
