package signup

import (
	"fmt"
	"github.com/5t4lk/blackjack/internal/ui"
	"github.com/5t4lk/blackjack/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func UserSignUp() error {
	var username, password string
	fmt.Print("[REGISTRATION] Please, enter your username: ")
	fmt.Scan(&username)
	fmt.Print("[REGISTRATION] Please, enter your password: ")
	fmt.Scan(&password)

	client, ctx, _, err := mongodb.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}

	err = mongodb.Ping(client, ctx)
	if err != nil {
		return err
	}

	var UserInfoBson interface{}

	UserInfoBson = bson.D{
		{username, password},
	}

	_, err = mongodb.InsertOne(client, ctx, "UsersDatabase", "UsersLogPass", UserInfoBson)
	if err != nil {
		return err
	}

	fmt.Print("[REGISTRATION] Successfully registered. Enjoy your gambling!\n")
	ui.Slow()
	return nil
}
