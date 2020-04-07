package main

import (
	"flag"
	"log"
	"time"

	"./services"
)

func main() {
	dir := flag.String("dir", "", "Configs directory")
	flag.Parse()
	if *dir == "" {
		log.Fatal("Directory is required")
	}
	services.ConfigDir = *dir
	err := services.Register()
	if err != nil {
		log.Panicln(err)
	}
	log.Println("Started service")

	// Run cron service
	for 1 == 1 {
		time.Sleep(60 * time.Second)
		services.ExecuteQueue()
	}
}
