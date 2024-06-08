package main

import (
	"fmt"
)

const DealerKey string = "dealer"

type Player struct {
	Hand     []Card
	Name     string
	isDealer bool
	isAi     bool
	Score    int
}

func (p Player) TotalHand() int {
	var sum int
	for i := range p.Hand {
		sum += p.Hand[i].GetRankValue()
	}

	return sum
}

func (p Player) PrintHand() {
	if p.isDealer {
		fmt.Println("Dealer Hand:")
	} else {
		fmt.Println("Current Hand:")
	}
	for i := range p.Hand {
		fmt.Printf("\t%s: %d\n", p.Hand[i], p.Hand[i].GetRankValue())
	}
	fmt.Printf("Total Hand: %d\n", p.TotalHand())
}

func (p Player) HasBlackjack() (has21 bool, blackjack bool) {

	var hasAce, hasFace bool

	for i := range p.Hand {
		if p.Hand[i].GetRankValue() == 10 {
			hasFace = true
		}

		if p.Hand[i].Rank == "ace" {
			hasAce = true
		}
	}

	return p.TotalHand() == blackjackLim, len(p.Hand) == 2 && hasAce && hasFace
}

func (p Player) IsBust() bool {
	return p.TotalHand() > blackjackLim
}

type Players map[string]Player

func (ps Players) ResetHands() {
	for k, v := range ps {
		v.Hand = []Card{}
		ps[k] = v
	}
}

func (ps Players) LoadHands(deck *[]Card) {
	for k, v := range ps {
		ps[k] = LoadHand(v, deck)
	}
}

func LoadPlayers() Players {
	dealer := Player{
		isDealer: true,
		Name:     "Dealer",
		Hand:     []Card{},
		Score:    0,
		isAi:     false,
	}

	players := map[string]Player{DealerKey: dealer}

	player := LoadPlayer()

	players[player.Name] = player

	return players
}

func LoadPlayer() Player {
	fmt.Print("Hello user please provide your name: ")

	var playerName string

	fmt.Scanln(&playerName)

	return Player{
		Hand:     []Card{},
		Name:     playerName,
		isAi:     false,
		isDealer: false,
		Score:    0,
	}
}
