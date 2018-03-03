package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := "*/3 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	go signalListen()

	select {}
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)
	for {
		<-c
		os.Exit(0)
	}
}
