package main

import (
	"fmt"
	"log"

	"github.com/MonSec/ctftime-discord-bot/api"
)

func main() {
	var monsec, err = api.TeamInfo(34111)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", monsec)
	upcomingEvents, err := api.UpcomingEvents(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", upcomingEvents)
}
