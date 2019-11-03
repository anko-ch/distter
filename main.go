package main

import (
	"log"
	"os"

	"./lib/pingpong"
	"github.com/bwmarrin/discordgo"
)

var (
	token   = "Bot " + os.Getenv("DISTTERBOT_TOKEN")
	stopper = make(chan bool)
)

func main() {
	session, err := discordgo.New()
	if err != nil {
		log.Fatalln(err)
	}
	session.Token = token

	packageHandlers := []interface{}{
		pingpong.OnMessage,
	}

	for _, handler := range packageHandlers {
		session.AddHandler(handler)
	}

	err = session.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()
	<-stopper
}
