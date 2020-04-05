package main;

import (
	"log"
	"flag"
	"services"
)

func main() {
   dir:=  flag.String("dir", "", "Configs directory")
   flag.Parse()
   if *dir == "" {
    log.Fatal("Directory is required")
   }
   log.Println("dir", *dir)
   services.Run(*dir)
   log.Println("Hello world")
}
