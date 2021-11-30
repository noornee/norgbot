package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/noornee/norgbot/config"
)




var BotID string


func main() {

	err := config.ReadConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(config.Token)

	dg, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Fatalf("Invalid bot parameters %v", err)
	}

	user, err := dg.User("@me")

	if err != nil {
		log.Fatal(err)
	}

	BotID = user.ID

	dg.AddHandler(MessageHandler)

	err = dg.Open()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Bot is running")

	// Cleanly close down the Discord session.
	<-make(chan struct{})

}



func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// fmt.Println(m.Content)

	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}

}
