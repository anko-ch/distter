package sendStatus

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessageSendChannelStatus(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "chan-stat" {
		session.ChannelMessageSend(message.ChannelID, channelStatus(session, message))
	}
}

func channelStatus(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var replyMessage string

	channel, err := session.Channel(
		message.ChannelID,
	)

	if err != nil {
		log.Println(err)
		replyMessage += err.Error()
	}

	replyMessage += fmt.Sprintf("```%v```", *channel)

	return replyMessage
}
