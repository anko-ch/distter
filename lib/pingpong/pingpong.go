package pingpong

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "ping" {
		_, err := session.ChannelMessageSend(message.ChannelID, "pong")
		if err != nil {
			log.Println(err)
		}
	}
}
