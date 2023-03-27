package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/1nv8rzim/Chandi-Bot/bot"
	"github.com/1nv8rzim/Chandi-Bot/commands"
)

func main() {

	commands.Init()

	err := bot.Session.Open()
	if err != nil {
		panic(err)
	}

	commands.StartScheduledTasks()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	commands.StopScheduledTasks()

	err = bot.Session.Close()
	if err != nil {
		panic(err)
	}

	log.Println("Bot stopped")
}
