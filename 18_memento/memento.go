package memento

import "fmt"

// 同一个类，用实例来保存状态
type Memento interface {}

type Game interface {
	Save() Memento
	Load(m Memento)
}

type gameMemento struct { // 小写开头，不允许外部访问， 可以由包内创建 // todo Memento大写开头
	hp, mp int // 想要保存的ConcreteGame的状态
}

type ConcreteGame struct {
	hp, mp int
}

func NewConcreteGame(hp, mp int) *ConcreteGame {
	return &ConcreteGame{
		hp:   hp,
		mp:   mp,
	}
}

func (g *ConcreteGame) Save() Memento {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *ConcreteGame) Load(m Memento) {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp
}

func (g *ConcreteGame) Play(hpDelta, mpDelta int) {
	g.hp += hpDelta
	g.mp += mpDelta
}

func (g *ConcreteGame) Status() {
	fmt.Printf("cur hp: %d, mp: %d\n", g.hp, g.mp)
}

