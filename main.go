package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var scorePlayer, scoreBot int

func main() {
	checkAnswer := playerEnter()
	if checkAnswer != true {
		log.Fatal("Unfortunately, you don't want to play.\n")
	}
	checkFirstCardPlayer, _ := dealFirstPlayerCard()
	if checkFirstCardPlayer != true {
		log.Fatal("error while dealing first player card\n")
	}

	checkFirstCardBot, _ := dealFirstBotCard()
	if checkFirstCardBot != true {
		log.Fatal("error while dealing first bot card")
	}

	for {
		askPlayer()
		if scorePlayer > 21 {
			fmt.Print("[GAME] You have lost. Score is more than 21.\n")
			slow()
			break
		} else if scorePlayer == 21 {
			fmt.Print("[GAME] You won. BLACKJACK!\n")
			slow()
			break
		}
	}
}

func playerEnter() bool {
	var playerChoice string
	fmt.Print("Do you want to play in `21`? Write [yes] or [no].\n")
	fmt.Scan(&playerChoice)

	if playerChoice == "yes" {
		fmt.Print("[GAME] You are started the game. Handing out cards!\n")
		slow()
		return true
	} else if playerChoice == "no" {
		return false
	}

	return false
}

func dealFirstBotCard() (bool, int) {
	firstBotCard := dealCard()
	showCardBot(firstBotCard)
	return true, scoreBot
}

func dealFirstPlayerCard() (bool, int) {
	firstPlayerCard := dealCard()
	showCardPlayer(firstPlayerCard)
	return true, scorePlayer
}

func askPlayer() bool {
	var choice int
	fmt.Printf("[GAME] Your score is %d.\n1. Take one more card\n2. Throw cards.\n", scorePlayer)
	fmt.Scan(&choice)
	slow()
	if choice == 1 {
		dealFirstPlayerCard()
	} else if choice == 2 {
		fmt.Print("[GAME] You threw the cards.\n")
		slow()
		botLogic()
		fmt.Printf("[GAME] Game is finished.\n")
		os.Exit(0)
	} else {
		fmt.Print("[ADMIN] Incorrect input.\n")
		os.Exit(0)
	}
	return true
}

func dealCard() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var card = rand.Intn(9)
	return card
}

func botLogic() {
	for i := 0; scoreBot < 14; i++ {
		takeCard, _ := dealFirstBotCard()
		if takeCard != true {
			log.Fatal("error while bot playing")
		}
	}
	grandFinal()
}

func grandFinal() {
	fmt.Printf("[GAME] BOT's score is %d.\n", scoreBot)
	if scorePlayer > scoreBot && scorePlayer <= 21 {
		slow()
		fmt.Printf("[GAME] You won! Enjoy!\n")
	} else if scorePlayer < scoreBot && scoreBot <= 21 {
		slow()
		fmt.Printf("[GAME] You lost! Next time be better!\n")
	} else if scorePlayer == scoreBot {
		slow()
		fmt.Printf("[GAME] It's a draw!")
	}
}

func showCardPlayer(player int) {
	switch player {
	case 0:
		fmt.Print("[GAME] You are dealt a card 6.\n")
		scorePlayer = 6 + scorePlayer
	case 1:
		fmt.Print("[GAME] You are dealt a card 7.\n")
		scorePlayer = 7 + scorePlayer
	case 2:
		fmt.Print("[GAME] You are dealt a card 8.\n")
		scorePlayer = 8 + scorePlayer
	case 3:
		fmt.Print("[GAME] You are dealt a card 9.\n")
		scorePlayer = 9 + scorePlayer
	case 4:
		fmt.Print("[GAME] You are dealt a card 10.\n")
		scorePlayer = 10 + scorePlayer
	case 5:
		fmt.Print("[GAME] You are dealt a card J.\n")
		scorePlayer = 2 + scorePlayer
	case 6:
		fmt.Print("[GAME] You are dealt a card Q.\n")
		scorePlayer = 3 + scorePlayer
	case 7:
		fmt.Print("[GAME] You are dealt a card K.\n")
		scorePlayer = 4 + scorePlayer
	case 8:
		fmt.Print("[GAME] You are dealt a card A.\n")
		scorePlayer = 11 + scorePlayer
	}
}
func showCardBot(bot int) {
	switch bot {
	case 0:
		fmt.Print("[GAME] BOT dealt a card 6.\n")
		scoreBot = 6 + scoreBot
	case 1:
		fmt.Print("[GAME] BOT dealt a card 7.\n")
		scoreBot = 7 + scoreBot
	case 2:
		fmt.Print("[GAME] BOT dealt a card 8.\n")
		scoreBot = 8 + scoreBot
	case 3:
		fmt.Print("[GAME] BOT dealt a card 9.\n")
		scoreBot = 9 + scoreBot
	case 4:
		fmt.Print("[GAME] BOT dealt a card 10.\n")
		scoreBot = 10 + scoreBot
	case 5:
		fmt.Print("[GAME] BOT dealt a card J.\n")
		scoreBot = 2 + scoreBot
	case 6:
		fmt.Print("[GAME] BOT dealt a card Q.\n")
		scoreBot = 3 + scoreBot
	case 7:
		fmt.Print("[GAME] BOT dealt a card K.\n")
		scoreBot = 4 + scoreBot
	case 8:
		fmt.Print("[GAME] BOT dealt a card A.\n")
		scoreBot = 11 + scoreBot
	}
}

func slow() {
	for i := 0; i < 3; i++ {
		time.Sleep(666 * time.Millisecond)
		fmt.Print("* ")
		if i == 2 {
			fmt.Print("\n")
		}
	}
}
