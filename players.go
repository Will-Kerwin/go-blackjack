package main

import "log/slog"

type Players map[string]Player

func (ps Players) ResetHands() {
	for k, v := range ps {
		v.Hand = Hand{}
		ps[k] = v
	}
}

func (ps Players) LoadHands(deck *Deck) {
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
	dealer := Player{
		isDealer: true,
		Name:     "Dealer",
		Hand:     []Card{},
		Score:    0,
		isAi:     false,
	}

	players := map[string]Player{DealerKey: dealer}

	player := CreatePlayer()

	players[player.Name] = player

	return players
}
