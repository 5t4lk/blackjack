package ux

import (
	"math/rand"
	"time"
)

type User struct {
	Balance float64
}

func (u *User) Won() {
	u.Balance = u.Balance + 5.00
}

func (u *User) Lost() {
	u.Balance = u.Balance - 5.00
}

func (u *User) Draw() {
	u.Balance = u.Balance - 0.75
}

func DealCard() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var card = rand.Intn(9)
	return card
}
