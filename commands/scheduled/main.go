package scheduled

import "github.com/bwmarrin/discordgo"

var (
	ComponentHandlers *map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
)

func Init(componentHandlers *map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)) {
	ComponentHandlers = componentHandlers
}
