package observer

import "fmt"

type Subscriber interface {
	Update(msg Msg)
}

type ConcreteSubscriber struct {
	name string
}

func NewConcreteSubscriber(name string) Subscriber {
	return &ConcreteSubscriber{name: name}
}

func (s *ConcreteSubscriber) Update(msg Msg) {
	fmt.Printf("sub: %s recv msg: %s \n", s.name, msg.Msg)
}