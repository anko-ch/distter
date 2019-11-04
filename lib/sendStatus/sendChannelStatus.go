package sendStatus

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnMessageSendChannelStatus(session *discordgo.Session, message *discordgo.MessageCreate) {
	switch message.Content {

	case "chan-stat":
		session.ChannelMessageSend(message.ChannelID, channelStatus(session, message))

	case "channel-permission":
		session.ChannelMessageSend(message.ChannelID, channelPermission(session, message))

	}
}

func channelStatus(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var replyMessage string

	channel, err := session.Channel(message.ChannelID)

	if err != nil {
		log.Println(err)
		replyMessage += err.Error()
	}

	replyMessage += fmt.Sprintf("%v", *channel)

	return fmt.Sprintf("```%s```", replyMessage)
}

func channelPermission(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var replyMessage string

	channel, err := session.Channel(message.ChannelID)

	if err != nil {
		log.Println(err)
		replyMessage += err.Error()
	}

	for _, permissionOverWrite := range channel.PermissionOverwrites {
		replyMessage += fmt.Sprintf("%v", permissionOverWrite)
	}

	return fmt.Sprintf("```%s```", replyMessage)
}
