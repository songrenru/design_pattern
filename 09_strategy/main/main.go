package main

import (
	"fmt"
	strategy "github.com/songrenru/design_attern/09_strategy"
)

func main() {
	p1 := strategy.NewPlayer("呆子", strategy.NewWinningStragety())
	p2 := strategy.NewPlayer("小机灵", strategy.NewProbStragety())

	for i := 0; i < 10000; i++ {
		h1 := p1.NextHand()
		h2 := p2.NextHand()
		if h1.IsStrongThan(h2) {
			p1.Win()
			p2.Lose()
		} else if h1.IsWeekThan(h2) {
			p1.Lose()
			p2.Win()
		} else {
			p1.Even()
			p2.Even()
		}
	}

	fmt.Println("result:")
	fmt.Println(p1)
	fmt.Println(p2)
}
