package commands

import (
	"github.com/1nv8rzim/Chandi-Bot/bot"
	"github.com/1nv8rzim/Chandi-Bot/commands/handlers"
	"github.com/1nv8rzim/Chandi-Bot/commands/scheduled"
	"github.com/1nv8rzim/Chandi-Bot/commands/slash"
	"github.com/1nv8rzim/Chandi-Bot/config"
	"github.com/1nv8rzim/Chandi-Bot/structs"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var (
	SlashCommands     map[string]*structs.SlashCommand                                      = make(map[string]*structs.SlashCommand)
	Handlers          map[string]interface{}                                                = make(map[string]interface{})
	ScheduledEvents   map[string]*structs.ScheduledEvent                                    = make(map[string]*structs.ScheduledEvent)
	ComponentHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) = make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate))

	quit chan interface{} = make(chan interface{})
)

func Init() {
	slash.Init(&ComponentHandlers)
	scheduled.Init(&ComponentHandlers)
	handlers.Init(&ComponentHandlers)

	populateSlashCommands()
	populateHandlers()
	populateScheduledEvents()

	addSlashCommands()
	addHandlers()

}

func addSlashCommands() {
	for _, command := range SlashCommands {
		_, err := bot.Session.ApplicationCommandCreate(config.AppID, config.GuildID, command.Command)
		if err != nil {
			logrus.WithError(err).Error("Failed to register slash command")
		} else {
			logrus.Infof("Registered slash command: %s", command.Command.Name)
		}
	}

	bot.Session.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {

		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			data := i.ApplicationCommandData()

			if command, ok := SlashCommands[data.Name]; ok {
				command.Handler(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if command, ok := ComponentHandlers[i.MessageComponentData().CustomID]; ok {
				command(s, i)
			}
		}
	})
}

func addHandlers() {
	for name, handler := range Handlers {
		bot.Session.AddHandler(handler)
		logrus.Infof("Registered handler: %s", name)
	}
}

func StartScheduledTasks() {
	for name, event := range ScheduledEvents {
		go event.Run(bot.Session, quit)
		logrus.Infof("Starting scheduled task: %v\n", name)
	}
}

func StopScheduledTasks() {
	if len(ScheduledEvents) > 0 {
		quit <- "kill"
	}
}
