package ux

import (
	"math/rand"
	"time"
)

func DealCard() int {
	time.Sleep(1 * time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	var card = rand.Intn(9)
	return card
}
