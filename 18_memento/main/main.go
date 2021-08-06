package main

import memento "github.com/songrenru/design_attern/18_memento"

func main() {
	g := memento.NewConcreteGame(100, 100)
	g.Status()

	g.Play(-20, -20)
	g.Status()
	m := g.Save()

	g.Play(-99, -99)
	g.Status()

	// 死翘翘了，复活一波
	g.Load(m)
	g.Status()

}
