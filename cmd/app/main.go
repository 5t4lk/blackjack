package main

import (
	"fmt"
	"github.com/5t4lk/blackjack/internal/ui"
	"github.com/5t4lk/blackjack/internal/ux"
	"log"
	"os"
)

var scorePlayer, scoreBot = 0, 0
var User = ux.User{Balance: 100.00}

func main() {
	if check := checkBalance(User.Balance); check != true {
		os.Exit(0)
	}

	scorePlayer, scoreBot = 0, 0

	fmt.Printf("[ADMIN] Your balance is %.2f$.\n", User.Balance)

	checkAnswer := ui.PlayerEnter()
	if checkAnswer != true {
		fmt.Print("[ADMIN] Unfortunately, you don't want to play. Table is closed.\n")
		os.Exit(0)
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
			User.Lost()
			ui.Slow()
			main()
		} else if scorePlayer == 21 {
			fmt.Print("[GAME] You won. BLACKJACK!\n")
			User.Won()
			ui.Slow()
			main()
		}
	}
}

func checkBalance(userBalance float64) bool {
	if userBalance == 0 || userBalance < 0 {
		fmt.Print("[ADMIN] You lost all money. Thanks for game.\n")
		return false
	}
	return true
}

func dealFirstBotCard() (bool, int) {
	firstBotCard := ux.DealCard()
	showCardBot(firstBotCard)
	return true, scoreBot
}

func dealFirstPlayerCard() (bool, int) {
	firstPlayerCard := ux.DealCard()
	showCardPlayer(firstPlayerCard)
	return true, scorePlayer
}

func askPlayer() bool {
	var choice int
	fmt.Printf("[GAME] Your score is %d.\n1. Take one more card\n2. Throw cards.\n", scorePlayer)
	fmt.Scan(&choice)

	if choice == 1 {
		dealFirstPlayerCard()
	} else if choice == 2 {
		fmt.Print("[GAME] You threw the cards.\n")
		ui.Slow()
		botLogic()
	} else {
		fmt.Print("[ADMIN] Incorrect input.\n")
		ui.Slow()
		main()
	}

	return true
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
		ui.Slow()
		fmt.Printf("[GAME] You won! Enjoy!\n")
		User.Won()
		main()
	} else if scorePlayer < scoreBot && scoreBot <= 21 {
		ui.Slow()
		fmt.Printf("[GAME] You lost! Next time be better!\n")
		User.Lost()
		main()
	} else if scorePlayer == scoreBot {
		ui.Slow()
		fmt.Printf("[GAME] It's a draw!\n")
		User.Draw()
		main()
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
