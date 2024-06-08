package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const blackjackLim int = 21

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Println("\nThank you for playing")
		os.Exit(1)
	}()

	fmt.Println("BlackJack!!!!")
	fmt.Println("\tBy Will Kerwin")

	players := LoadPlayers()
	defer players.PrintScores()

	for {
		menuOption := Menu()

		switch menuOption {
		case 0:
			return
		case 1:
			var rounds int
			fmt.Print("How may rounds do you want to play?: ")
			fmt.Scanln(&rounds)
			PlayGame(players, rounds)
		default:
			fmt.Println("Invalid Option")
		}
	}
}

func PlayGame(players Players, rounds int) {
	deck := LoadDeck()

	fmt.Printf("Playing %d rounds of blackjack\n", rounds)

	// run for total rounds provided
	for i := 0; i < rounds; i++ {
		fmt.Printf("\n\n***Round %d***\n\n", i+1)
		var roundComplete bool

		players.LoadHands(&deck)

		// loop round
		for !roundComplete {
			dealer := players[DealerKey]

			dealer.PrintHand()
			fmt.Println()

			dealer.IsInitiallyBust()
			players[DealerKey] = dealer

			for k, v := range players {
				if v.isDealer {
					continue
				}
				roundComplete = v.PlayHand(&deck)
				players[k] = v
			}

			dealer.PlayDealerHand(&deck)
			players[DealerKey] = dealer

			if !dealer.IsBust() {
				for k, v := range players {
					if v.isDealer {
						continue
					}
					if v.TotalHand() >= dealer.TotalHand() && !v.IsBust() {
						v.Score += 1
						players[k] = v
					} else {
						if d21, _ := dealer.HasBlackjack(); d21 || dealer.IsBust() {
							continue
						}
						dealer.Score += 1
						players[DealerKey] = dealer
					}
				}
			}

		}

		players.ResetHands()
	}

	players.PrintScores()
}

func Menu() int {
	var option int
	fmt.Println("***Menu***")
	fmt.Println("\t1: New Game")
	fmt.Println("\t0: Exit")
	fmt.Print("Please Select An Option: ")
	fmt.Scanln(&option)
	return option
}
