package helpers

import (
	"github.com/bwmarrin/discordgo"
)

func JumpURL(m *discordgo.Message) string {
	return "https://discordapp.com/channels/" + config.GuildID + "/" + m.ChannelID + "/" + m.ID
}
