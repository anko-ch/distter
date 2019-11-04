package personalChannel

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const defaultRoleColor = 0xffffff

func OnJoin(session *discordgo.Session, newUser *discordgo.GuildMemberAdd) {

	readWriteRoleName := "ReadWrite:" + newUser.User.ID
	roles, err := session.GuildRoles(newUser.GuildID)

	// 重複してた場合処理を中断する
	for _, role := range roles {
		if role.Name == readWriteRoleName {
			return
		}
	}

	ReadWriteRole, err := GuildRoleCreateEdit(
		session,
		newUser.GuildID,
		readWriteRoleName,
	)

	err = session.GuildMemberRoleAdd(newUser.GuildID, newUser.User.ID, ReadWriteRole.ID)

	_, err = session.GuildChannelCreate(
		newUser.GuildID,
		newUser.User.Username,
		discordgo.ChannelTypeGuildText,
	)

	if err != nil {
		log.Println(err)
	}

}

func OnMessageSendChannelStatus(session *discordgo.Session, message discordgo.MessageCreate) {
	if message.Content[0:9] == "chan-stat" {
		channel, err := session.Channel(
			strings.Replace(message.Content, "chan-stat", "", 1),
		)
		if err != nil {
			log.Println(err)
			session.ChannelMessageSend(message.ChannelID, err.Error())
		}

		session.ChannelMessageSend(
			message.ChannelID,
			fmt.Sprintf("%v", *channel),
		)
	}
}

//code review: この関数必要ですかね?同じようなことを二回するので、一応分けて書いたのですが...
func GuildRoleCreateEdit(session *discordgo.Session, guildID, roleName string) (*discordgo.Role, error) {
	role, err := session.GuildRoleCreate(guildID)
	if err != nil {
		return role, err
	}

	role, err = session.GuildRoleEdit(guildID, role.ID, roleName, defaultRoleColor, false, 0, false)
	if err != nil {
		return role, err
	}

	return role, nil
}
