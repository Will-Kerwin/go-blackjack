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

	for i := 0; i < rounds; i++ {
		fmt.Printf("Round %d\n\n", i+1)
		var roundComplete bool

		players.LoadHands(&deck)

		for !roundComplete {
			dealer := players[DealerKey]

			dealer.PrintHand()
			fmt.Println()

			if has21, blackjack := players[DealerKey].HasBlackjack(); has21 {
				dealer.PrintHand()
				dealer.Score += 1
				if blackjack {
					dealer.Score += 2
					fmt.Println("Dealer has blackjack")
				} else {
					fmt.Println("Dealer Has 21")
				}

				players[DealerKey] = dealer
			} else if players[DealerKey].IsBust() {
				fmt.Println("Dealer Bust")
			}

			for k, v := range players {
				if v.isDealer {
					continue
				}
				var option int
				for {
					v.PrintHand()
					if has21, blackjack := v.HasBlackjack(); has21 {
						v.Score += 1
						if blackjack {
							v.Score += 2
							fmt.Printf("%s has blackjack\n", k)
						} else {
							fmt.Printf("%s Has 21\n", k)
						}

						players[k] = v
						roundComplete = true
						break
					} else if v.IsBust() {
						fmt.Println("You are bust")
						roundComplete = true
						break
					}

					fmt.Print("Please enter 0 for stick and 1 for hit: ")
					fmt.Scanln(&option)

					if option == 0 {
						break
					}

					v.Hand = append(v.Hand, Draw(&deck))
					players[k] = v
				}
			}
		}

		players.ResetHands()
	}
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
