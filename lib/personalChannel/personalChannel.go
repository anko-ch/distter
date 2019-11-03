package personalChannel

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func OnJoin(session *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	_, err := session.GuildChannelCreate(
		newMember.GuildID,
		newMember.User.Username,
		discordgo.ChannelTypeGuildText,
	)
	if err != nil {
		log.Println(err)
	}

}
