package part1

import (
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	card string
	bid  int
}

type kv struct {
	Key   byte
	Value int
}

type CardType int64

const (
	HighCard CardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (h Hand) FrequencySortedByValue() []kv {
	freqs := h.Frequency()

	var ss []kv
	for k, v := range freqs {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value // sort in descending order
	})

	return ss
}
func (h Hand) Frequency() map[byte]int {
	freqs := make(map[byte]int)
	for i := 0; i < len(h.card); i++ {
		freqs[h.card[i]]++
	}

	return freqs
}

func (h Hand) IsFiveOfAKind() bool {
	return len(h.Frequency()) == 1
}

func (h Hand) IsFourOfAKind() bool {
	return len(h.Frequency()) == 2 && h.FrequencySortedByValue()[0].Value == 4
}

func (h Hand) IsFullHouse() bool {
	return len(h.Frequency()) == 2 && h.FrequencySortedByValue()[0].Value == 3
}

func (h Hand) IsThreeOfAKind() bool {
	return len(h.Frequency()) == 3 && h.FrequencySortedByValue()[0].Value == 3
}

func (h Hand) IsTwoPair() bool {
	return len(h.Frequency()) == 3 && h.FrequencySortedByValue()[0].Value == 2
}

func (h Hand) IsOnePair() bool {
	return len(h.Frequency()) == 4 && h.FrequencySortedByValue()[0].Value == 2
}

func (h Hand) IsHighCard() bool {
	return len(h.Frequency()) == 5
}

func (h Hand) Type() CardType {
	if h.IsFiveOfAKind() {
		return FiveOfAKind
	} else if h.IsFourOfAKind() {
		return FourOfAKind
	} else if h.IsFullHouse() {
		return FullHouse
	} else if h.IsThreeOfAKind() {
		return ThreeOfAKind
	} else if h.IsTwoPair() {
		return TwoPair
	} else if h.IsOnePair() {
		return OnePair
	} else if h.IsHighCard() {
		return HighCard
	} else {
		panic("unknown type")
	}
}

func CompareCard(a, b byte) bool {
	prio := []byte{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	for i := 0; i < len(prio); i++ {
		if a == prio[i] {
			return true
		} else if b == prio[i] {
			return false
		} else {
			continue
		}
	}
	panic("unknown state")
}

func Greater(a, b Hand) bool {
	if a.Type() == b.Type() {
		for i := 0; i < len(a.card); i++ {
			if a.card[i] == b.card[i] {
				continue
			} else {
				return CompareCard(a.card[i], b.card[i])
			}
		}
	}
	return a.Type() > b.Type()

}

func Solve(input string) int {
	sum := 0
	hands := make([]Hand, 0)

	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{split[0], bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		return Greater(hands[j], hands[i])
	})

	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	return sum
}
