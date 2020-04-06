package main

import (
	"flag"
	"log"

	"./services"
)

func main() {
	dir := flag.String("dir", "", "Configs directory")
	flag.Parse()
	if *dir == "" {
		log.Fatal("Directory is required")
	}
	services.Register(*dir)
	log.Println("Started service")

	// Run cron service
	// gocron.Every(1).Minute().Do(services.ExecuteQueue)
	// <-gocron.Start()
}
