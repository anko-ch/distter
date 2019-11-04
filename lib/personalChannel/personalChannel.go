package personalChannel

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

const defaultRoleColor = 0xffffff

func OnJoin(session *discordgo.Session, newMember *discordgo.GuildMemberAdd) {

	ReadWriteRole, err := GuildRoleCreateEdit(
		session,
		newMember.GuildID,
		"ReadWrite:"+newMember.User.ID,
	)

	_, err = GuildRoleCreateEdit(
		session,
		newMember.GuildID,
		"Read:"+newMember.User.ID,
	)

	err = session.GuildMemberRoleAdd(newMember.GuildID, newMember.User.ID, ReadWriteRole.ID)

	_, err = session.GuildChannelCreate(
		newMember.GuildID,
		newMember.User.Username,
		discordgo.ChannelTypeGuildText,
	)

	if err != nil {
		log.Println(err)
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
