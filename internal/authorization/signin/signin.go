package signin

import (
	"fmt"
	"github.com/5t4lk/blackjack/internal/ui"
	"github.com/5t4lk/blackjack/pkg/db/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func UserSignIn() error {
	var username, password string
	fmt.Print("[AUTHORIZATION] Please, enter your username: ")
	fmt.Scan(&username)
	fmt.Print("[AUTHORIZATION] Please, enter your password: ")
	fmt.Scan(&password)

	client, ctx, _, err := mongodb.Connect("mongodb://localhost:27017")
	if err != nil {
		return err
	}

	err = mongodb.Ping(client, ctx)
	if err != nil {
		return err
	}

	var filter, option interface{}

	filter = bson.D{
		{username, password},
	}

	cursor, err := mongodb.Query(client, ctx, "UsersDatabase", "UsersLogPass", filter, option)
	if err != nil {
		return err
	}

	var results []bson.D

	if err = cursor.All(ctx, &results); err != nil {
		return err
	}
	if results == nil {
		fmt.Print("[AUTHORIZATION] User is not found. Please make sure the input is correct.\n")
		ui.Slow()
		UserAuthorization()
	} else {
		fmt.Printf("[AUTHORIZATION] Welcome %s! Enjoy your playing!\n", username)
		ui.Slow()
		return nil
	}

	return nil
}
