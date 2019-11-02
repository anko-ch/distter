package main

import (
	"github.com/bwmarrin/discordgo"
	"./lib/pingpong"
	"log"
)

var (
	token = "Bot ***"
	stopBot = make(chan bool)
	packageHandlers = []interface{}{
		pingpong.OnMessage,
	}
)

func main() {
	session, err := discordgo.New()
	session.Token = token
	if err != nil{
		log.Fatalln(err)
	}

	for _, handler := range packageHandlers{
		session.AddHandler(handler)
	}

	err = start(session)
	if err != nil{
		log.Fatalln(err)
	}
}

func start(session *discordgo.Session)(error){
	err := session.Open()
	if err != nil{
		return err
	}
	<- stopBot
	return nil
}