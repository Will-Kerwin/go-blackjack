package card

import (
	"fmt"
	"testing"
)

func TestGetRankValueTableDriven(t *testing.T) {

	var tests = []struct {
		card Card
		want int
	}{
		{
			card: Card{
				Rank: "queen",
				Suit: "hearts",
			},
			want: 10,
		},
		{
			card: Card{
				Rank: "ace",
				Suit: "diamonds",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Get Rank Value for %s of %s", tt.card.Rank, tt.card.Suit)
		t.Run(testname, func(t *testing.T) {
			ans := tt.card.GetRankValue()
			if ans != tt.want {
				t.Errorf("got value of %d, want %d", ans, tt.want)
			}
		})
	}

}

func TestStringTableDriven(t *testing.T) {
	var tests = []struct {
		card Card
		want string
	}{
		{
			card: Card{
				Rank: "queen",
				Suit: "hearts",
			},
			want: "queen of hearts",
		},
		{
			card: Card{
				Rank: "ace",
				Suit: "diamonds",
			},
			want: "ace of diamonds",
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Get string name for card %s of %s", tt.card.Rank, tt.card.Suit)

		t.Run(testname, func(t *testing.T) {
			ans := tt.card.String()
			if ans != tt.want {
				t.Errorf("got string value of %s, want %s", ans, tt.want)
			}
		})
	}
}
