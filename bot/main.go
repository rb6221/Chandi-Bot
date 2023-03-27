package bot

import (
	"github.com/1nv8rzim/Chandi-Bot/config"
	"github.com/bwmarrin/discordgo"
)

var (
	Session *discordgo.Session
)

func init() {
	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		panic(err)
	}

	Session = session

	Session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAll)
}
