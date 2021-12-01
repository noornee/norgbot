package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/noornee/norgbot/config"
)

var (
	BotID string
	goBot *discordgo.Session
	err   error
)

func Start() {

	goBot, err = discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Fatalf("Invalid bot parameters %v", err)
	}

	user, err := goBot.User("@me")

	if err != nil {
		log.Fatal(err)
	}

	BotID = user.ID

	goBot.AddHandler(MessageHandler)

	err = goBot.Open()


	if err != nil {
		log.Fatalf("There was an error launching bot %v", err)
	}

	fmt.Println("Bot is running")

}

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	echo := config.BotPrefix + "echo "

	if strings.HasPrefix(m.Content, echo) {
		if m.Author.ID == BotID {
			return
		}
		s.ChannelMessageSend(m.ChannelID, m.Content[len(echo):])
	}

}
