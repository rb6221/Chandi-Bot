package structs

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type ScheduledEvent struct {
	Event func(*discordgo.Session, chan interface{}) error
}

func NewScheduledTask(function func(*discordgo.Session, chan interface{}) error) *ScheduledEvent {
	return &ScheduledEvent{
		Event: function,
	}
}

func (e *ScheduledEvent) Run(s *discordgo.Session, quit chan interface{}) {
	for {
		err := e.Event(s, quit)
		if err != nil {
			log.Println(err)
		} else {
			return
		}
	}
}
