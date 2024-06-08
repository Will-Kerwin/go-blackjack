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

func (d *Deck) shuffleDeck() {
	// shuffle a deck of cards randomly based on length
	slog.Debug("Shuffling Deck")
	deck := *d
	for i := range deck {
		j := rand.Intn(len(deck) - 1)
		deck[i], deck[j] = deck[j], deck[i]
	}

	*d = deck
}

func LoadDeck() Deck {
	// load a shuffled deck

	deck := generateDeck()

	deck.shuffleDeck()

	return deck
}

func (deckPtr *Deck) Draw() Card {
	slog.Debug("Drawing a card")
	var card Card
	deck := *deckPtr
	// draw a card from a deck checking if it is the last card
	if deck.IsLastCard() {
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

func (deck Deck) IsLastCard() bool {
	return len(deck) == 1
}
