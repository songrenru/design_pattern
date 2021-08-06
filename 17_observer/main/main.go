package main

import (
	observer "github.com/songrenru/design_attern/17_observer"
	"strconv"
)

func main() {
	pub := observer.NewConcretePublisher()
	sub1 := observer.NewConcreteSubscriber("sub1")
	sub2 := observer.NewConcreteSubscriber("sub2")
	sub3 := observer.NewConcreteSubscriber("sub3")
	sub4 := observer.NewConcreteSubscriber("sub4")
	sub5 := observer.NewConcreteSubscriber("sub5")
	pub.AddSubscriber(sub1)
	pub.AddSubscriber(sub2)
	pub.AddSubscriber(sub3)
	pub.AddSubscriber(sub4)
	pub.AddSubscriber(sub5)

	for i := 0; i < 10; i++ {
		msg :=observer.Msg{
			Id:  i,
			Msg: "msg:" + strconv.Itoa(i),
		}
		pub.Notify(msg)
	}
}
