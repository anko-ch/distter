package personalChannel

import (
	"log"

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
