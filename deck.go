package main

import (
	"log/slog"
	"math/rand"
)

type Deck []Card

func generateDeck() Deck {
	// Generate a deck of cards from rank and suits
	slog.Debug("Generating Deck")
	deck := make(Deck, 0)

	for rK := range Ranks {
		for sK := range Suits {
			card := Card{
				Suit: Suits[sK],
				Rank: rK,
			}

			deck = append(deck, card)
		}
	}

	slog.Debug("Generated deck of cards", "len", len(deck))

	return deck
}

func shuffleDeck(c Deck) {
	// shuffle a deck of cards randomly based on length
	slog.Debug("Shuffling Deck")
	for i := range c {
		j := rand.Intn(len(c) - 1)
		c[i], c[j] = c[j], c[i]
	}
}

func LoadDeck() Deck {
	// load a shuffled deck

	deck := generateDeck()

	shuffleDeck(deck)

	return deck
}

func Draw(deckPtr *Deck) Card {
	slog.Debug("Drawing a card")
	var card Card
	deck := *deckPtr
	// draw a card from a deck checking if it is the last card
	if IsLastCard(deck) {
		card = deck[0]
		slog.Debug("Drew", "card", card)
		*deckPtr = LoadDeck()
		return card
	}

	card = deck[0]
	slog.Debug("Drew", "card", card)
	*deckPtr = deck[1:]
	return card
}

func IsLastCard(deck Deck) bool {
	return len(deck) == 1
}

func LoadHand(player Player, deck *Deck) Player {
	slog.Debug("Loading hand for player", "name", player.Name)
	for i := 0; i < 2; i++ {
		player.Hand = append(player.Hand, Draw(deck))
	}

	slog.Debug("Player has hand of cards", "name", player.Name, "handLen", len(player.Hand))

	return player
}
