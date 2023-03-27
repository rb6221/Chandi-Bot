package structs

import "github.com/bwmarrin/discordgo"

type SlashCommand struct {
	Command *discordgo.ApplicationCommand
	Handler func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func NewSlashCommand(command *discordgo.ApplicationCommand, handler func(s *discordgo.Session, i *discordgo.InteractionCreate)) *SlashCommand {
	return &SlashCommand{
		Command: command,
		Handler: handler,
	}
}
