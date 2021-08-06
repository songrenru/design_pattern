package strategy

import (
	"math/rand"
	"time"
)

// 策略的抽象
type Stragety interface {
	NextHand() *Hand
	Study(win bool)
}

// 具体策略WinningStragety
type WinningStragety struct {
	win bool
	preHand *Hand
}

func NewWinningStragety() Stragety {
	s := &WinningStragety{}
	s.preHand = Hands[getRandHandValue(3)]

	return s
}

func (s *WinningStragety) NextHand() *Hand {
	if s.win {
		return s.preHand
	}
	return GetHand(getRandHandValue(3))
}

func (s *WinningStragety) Study(win bool) {
	s.win = win
}

func getRandHandValue(n int) int {
	if n == 0 {
		return 0
	}
	rand.Seed(int64(time.Now().Second()))
	return rand.Intn(n)
}


// 具体策略ProbStragety
type ProbStragety struct {
	win bool
	preHandVal int
	currHandVal int
	history [3][3]int
}

func NewProbStragety() Stragety {
	s := &ProbStragety{
		history: [3][3]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
		preHandVal: getRandHandValue(3),
		currHandVal: getRandHandValue(3),
	}
	return s
}

func (s *ProbStragety) NextHand() *Hand {
	bet := getRandHandValue(s.getSum(s.currHandVal))
	handVal := 2
	if bet < s.history[s.currHandVal][0] {
		handVal = 0
	} else if bet < s.history[s.currHandVal][0] + s.history[s.currHandVal][0] {
		handVal = 1
	}
	s.preHandVal = s.currHandVal
	s.currHandVal = handVal
	return GetHand(handVal)
}

func (s *ProbStragety) Study(win bool) {
	if win {
		s.history[s.preHandVal][(s.currHandVal + 1) % 3]++
	} else {
		s.history[s.preHandVal][(s.currHandVal + 2) % 3]++
	}
}

func (s *ProbStragety) getSum(hv int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		sum += s.history[hv][i]
	}
	return sum
}
