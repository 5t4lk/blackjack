package ux

import (
	"fmt"
	"github.com/5t4lk/blackjack/internal/ui"
	"math/rand"
	"time"
)

type User struct {
	Bet     float64
	Deposit float64
}

func (u *User) Won() {
	u.Deposit = u.Deposit + u.Bet
}

func (u *User) Lost() {
	u.Deposit = u.Deposit - u.Bet
}

func (u *User) Draw() {
	u.Deposit = u.Deposit - 0.75
}

func (u *User) AskDeposit() {
	fmt.Print("[ADMIN] Enter your deposit: \n")
	fmt.Scan(&u.Deposit)
	if u.Deposit < 0 {
		fmt.Print("[ADMIN] Enter correct deposit.\n")
		u.AskDeposit()
	}
}

func (u *User) AskBet() {
	fmt.Print("[ADMIN] Enter your bet: \n")
	fmt.Scan(&u.Bet)
	if u.Bet <= 0 {
		fmt.Print("[ADMIN] Enter correct bet.\n")
		ui.Slow()
		u.AskBet()
	} else if u.Bet > u.Deposit {
		fmt.Print("[ADMIN] Bet more than your balance.\n")
		ui.Slow()
		u.AskBet()
	}
}

func DealCard() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var card = rand.Intn(9)
	return card
}
