package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gocolly/colly"
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
	anime := config.BotPrefix + "anime "

	ch := make(chan string) // channel for the anime url

	if m.Author.ID == BotID {
		return
	}

	if strings.HasPrefix(m.Content, echo) {
		s.ChannelMessageSend(m.ChannelID, m.Content[len(echo):])
	}

	if strings.HasPrefix(m.Content, anime) {
		word := m.Content[len(anime):]
		go AnimeScraper(word, ch)
		for {
			s.ChannelMessageSend(m.ChannelID, <-ch)
		}

	}

}

// A function that makes use of the colly framework to scrape data from an anime site.
func AnimeScraper(word string, ch chan string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www1.gogoanime.cm"),
	)

	c.OnHTML("div.img a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		ch <- r.URL.String()
	})

	c.Visit(fmt.Sprintf("https://www1.gogoanime.cm//search.html?keyword=%v", word))

}
