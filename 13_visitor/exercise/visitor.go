package visitor

// Visitor
type Visitor interface {
	Visit(e Element)
}

// acceptor
type Element interface {
	Accept(v Visitor)
}
