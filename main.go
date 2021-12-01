package main

import (
	"log"

	"github.com/noornee/norgbot/bot"
	"github.com/noornee/norgbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	bot.Start()

	<-make(chan struct{})

}
