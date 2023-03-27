package helpers

import (
	"github.com/1nv8rzim/Chandi-Bot/config"
	"github.com/bwmarrin/discordgo"
)

func JumpURL(m *discordgo.Message) string {
	return "https://discordapp.com/channels/" + config.GuildID + "/" + m.ChannelID + "/" + m.ID
}
