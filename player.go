package main

import (
	"fmt"
	"log/slog"
)

const DealerKey string = "dealer"
const DealerStickLim int = 17

type Player struct {
	Hand     Hand
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
	fmt.Println()
	if p.isDealer {
		fmt.Println("Dealer Hand:")
		if p.IsAtDealerStickLim() {
			fmt.Println("Dealer will stick!")
		}
	} else {
		fmt.Println("Current Hand:")
	}
	for i := range p.Hand {
		fmt.Printf("\t%s: %d\n", p.Hand[i], p.Hand[i].GetRankValue())
	}
	fmt.Printf("Total Hand: %d\n", p.TotalHand())
}

func (p Player) PrintScore() {
	if p.isDealer {
		fmt.Printf("Dealer Score: %d\n", p.Score)
	} else {
		fmt.Printf("%s Score: %d\n", p.Name, p.Score)
	}
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

func (p *Player) LoadHand(deck *Deck) {
	slog.Debug("Loading hand for player", "name", p.Name)
	for i := 0; i < 2; i++ {
		p.Hand = append(p.Hand, deck.Draw())
	}

	slog.Debug("Player has hand of cards", "name", p.Name, "handLen", len(p.Hand))
}

func CreatePlayer() Player {
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

func (p *Player) PlayHand(deck *Deck) bool {
	var option int
	for {
		p.PrintHand()
		if has21, blackjack := p.HasBlackjack(); has21 {
			p.Score += 1
			if blackjack {
				p.Score += 2
				fmt.Printf("%s has blackjack\n", p.Name)
			} else {
				fmt.Printf("%s Has 21\n", p.Name)
			}

			return true
		} else if p.IsBust() {
			fmt.Println("You are bust")
			return true
		}

		fmt.Print("Please enter 0 for stick and 1 for hit: ")
		fmt.Scanln(&option)

		if option == 0 {
			return true
		}

		p.Hand = append(p.Hand, deck.Draw())
	}
}

func (p Player) IsAtDealerStickLim() bool {
	return p.TotalHand() > DealerStickLim
}

func (p *Player) PlayDealerHand(deck *Deck) {
	for {
		p.PrintHand()
		if has21, blackjack := p.HasBlackjack(); has21 {
			p.Score += 1
			if blackjack {
				p.Score += 2
				fmt.Printf("%s has blackjack\n", p.Name)
			} else {
				fmt.Printf("%s Has 21\n", p.Name)
			}

			return
		} else if p.IsBust() {
			fmt.Println("Dealer is bust")
			return
		}

		if p.IsAtDealerStickLim() {
			return
		}

		p.Hand = append(p.Hand, deck.Draw())
	}
}

func (p *Player) IsInitiallyBust() {
	if has21, blackjack := p.HasBlackjack(); has21 {
		p.PrintHand()
		p.Score += 1
		if blackjack {
			p.Score += 2
			fmt.Println("Dealer has blackjack")
		} else {
			fmt.Println("Dealer Has 21")
		}
	} else if p.IsBust() {
		fmt.Println("Dealer Bust")
	}
}
