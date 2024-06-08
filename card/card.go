package card

import (
	"fmt"
)

type Card struct {
	Suit string
	Rank string
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}

func (c Card) GetRankValue() int {
	return Ranks[c.Rank]
}

type Hand []Card

var Suits []string = []string{"hearts", "clubs", "spades", "diamonds"}

var Ranks map[string]int = map[string]int{
	"ace":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}
