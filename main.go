package main

import (
	"log"

	"./lib/pingpong"
	"github.com/bwmarrin/discordgo"
)

const token = "Bot ***"

var stopper = make(chan bool)

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
