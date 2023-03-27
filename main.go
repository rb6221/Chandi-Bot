package main

import (
	"log"
	"os"
	"os/signal"
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
	data.Close()

	err = bot.Session.Close()
	if err != nil {
		panic(err)
	}

	log.Println("Bot stopped")
}
