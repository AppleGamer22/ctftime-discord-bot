package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/MonSec/ctftime-discord-bot/bot"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var session *discordgo.Session

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	session, err = discordgo.New(os.Getenv("TOKEN"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	err := session.Open()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()
	log.Println("Connected to Discord!")
	_, err = session.ApplicationCommandCreate(os.Getenv("APP_ID"), "", bot.TeamCommand)
	if err != nil {
		log.Println(err)
		return
	}
	session.AddHandler(bot.TeamCommandHandler)
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down..")
	// var monsec, err = api.TeamInfo("34111")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", monsec)
	// upcomingEvents, err := api.UpcomingEvents(nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", upcomingEvents)

}
