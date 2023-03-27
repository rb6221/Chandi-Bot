package structs

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
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
			logrus.Println(err)
		} else {
			return
		}
	}
}
