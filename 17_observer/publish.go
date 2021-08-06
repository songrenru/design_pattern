package observer

type Publisher interface {
	AddSubscriber(s Subscriber)
	RemoveSubscriber(s Subscriber)
	Notify(msg Msg)
}

type Msg struct {
	Id int
	Msg string 
}

type ConcretePublisher struct {
	ss map[Subscriber]struct{} // todo map存索引，slice存sub,保证遍历顺序
}

func NewConcretePublisher() Publisher {
	return &ConcretePublisher{
		ss: make(map[Subscriber]struct{}),
	}
}

func (p *ConcretePublisher) AddSubscriber(s Subscriber) {
	p.ss[s] = struct{}{}
}

func (p *ConcretePublisher) RemoveSubscriber(s Subscriber) {
	delete(p.ss, s)
}

func (p *ConcretePublisher) Notify(msg Msg) {
	for s := range p.ss {
		s.Update(msg)
	}
}