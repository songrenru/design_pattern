package strategy

import "fmt"

type Player struct {
	name string
	strategy Stragety
	winCount, loseCount, gameCount int
}

func NewPlayer(name string, s Stragety) *Player {
	return &Player{
		name: name,
		strategy: s,
	}
}

func (p *Player) NextHand() *Hand {
	return p.strategy.NextHand()
}

func (p *Player) Study(win bool) {
	p.strategy.Study(win)
}

func (p *Player) Win() {
	p.strategy.Study(true)
	p.winCount++
	p.gameCount++
}

func (p *Player) Lose() {
	p.strategy.Study(false)
	p.loseCount++
	p.gameCount++
}

func (p *Player) Even() {
	p.strategy.Study(false)
	p.gameCount++
}

func (p *Player) String() string {
	return fmt.Sprintf("name: %s, total: %d, win: %d, lose: %d", p.name, p.gameCount, p.winCount, p.loseCount)
}


