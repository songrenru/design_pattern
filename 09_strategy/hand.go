package strategy

var (
	HandStone = &Hand{name: "石头", value: 0}
	HandJiandao = &Hand{name: "剪刀", value: 1}
	HandBu = &Hand{name: "布", value: 2}
	Hands = [3]*Hand{HandStone, HandJiandao, HandBu}
)

func GetHand(handValue int) *Hand {
	return Hands[handValue]
}

type Hand struct {
	name string
	value int
}

func (h *Hand) IsStrongThan(hand *Hand) bool {
	return h.fight(hand) == 1
}

func (h *Hand) IsWeekThan(hand *Hand) bool {
	return h.fight(hand) == -1
}

func (h *Hand) fight(hand *Hand) int {
	if h == hand {
		return 0
	} else if (h.value+1)%3 == hand.value {
		return 1;
	} else {
		return -1
	}
}

func (h *Hand) String() string {
	return h.name
}


