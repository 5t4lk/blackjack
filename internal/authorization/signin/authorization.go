package signin

import (
	"fmt"
	"github.com/5t4lk/blackjack/internal/authorization/signup"
	"github.com/5t4lk/blackjack/internal/ui"
)

func UserAuthorization() error {
	var userAsk int
	fmt.Print("[ADMIN] Welcome to our application. To continue playing, you need to log in.\n" +
		"[ADMIN] Are you already have a profile and want to sign-in or you don't and\n[ADMIN] want to sign-up?\n[1] - Sign-in\n[2] - Sign-up\n")
	fmt.Scan(&userAsk)

	if userAsk == 1 {
		err := UserSignIn()
		if err != nil {
			return err
		}
	} else if userAsk == 2 {
		err := signup.UserSignUp()
		if err != nil {
			return err
		}
	} else {
		fmt.Print("[APP] Incorrect input. Try again!\n")
		ui.Slow()
		UserAuthorization()
	}

	return nil
}
