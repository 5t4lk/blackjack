package ui

import (
	"fmt"
	"os"
	"time"
)

func PlayerEnter() bool {
	var playerChoice string

	fmt.Print("[ADMIN] Do you want to play in `21`? Write [yes] or [no].\n")
	fmt.Scan(&playerChoice)

	if playerChoice == "yes" {
		fmt.Print("[GAME] You are started the game. Handing out cards!\n")
		Slow()
		return true
	} else if playerChoice == "no" {
		return false
	} else {
		fmt.Print("[ADMIN] Incorrect input.\n")
		os.Exit(0)
	}

	return false
}

func Slow() {
	for i := 0; i < 3; i++ {
		time.Sleep(666 * time.Millisecond)
		fmt.Print("* ")
		if i == 2 {
			fmt.Print("\n")
		}
	}
}
