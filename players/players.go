package players

import (
	c "blackjack/card"
	d "blackjack/deck"
	p "blackjack/player"
	"log/slog"
)

type Players map[string]p.Player

func (ps Players) ResetHands() {
	for k, v := range ps {
		v.Hand = c.Hand{}
		ps[k] = v
	}
}

func (ps Players) LoadHands(deck *d.Deck) {
	for k, v := range ps {
		v.LoadHand(deck)
		slog.Debug("player has hand", "hand", v.Hand)
		ps[k] = v
	}
}

func (ps Players) PrintScores() {
	for _, v := range ps {
		v.PrintScore()
	}
}

func LoadPlayers() Players {
	dealer := p.Player{
		IsDealer: true,
		Name:     "Dealer",
		Hand:     []c.Card{},
		Score:    0,
		IsAi:     false,
	}

	players := map[string]p.Player{}
	players[p.DealerKey] = dealer

	player := p.CreatePlayer()

	players[player.Name] = player

	return players
}
