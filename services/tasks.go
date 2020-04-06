package services

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"
)

// ExecuteQueue read queue file and execute based on defination
func ExecuteQueue() error {
	// Read schedule file
	file, err := os.Open("schedule.log")
	if err != nil {
		return err
	}
	defer file.Close()
	var scheduleItems string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scheduleItems = scanner.Text()
	}
	// Get task to execute next
	// Check in the current hour minute
	// Schedule execution asyncronously
	var fileResult []map[string]string
	err = json.Unmarshal([]byte(scheduleItems), &fileResult)
	if err != nil {
		log.Panicln(err)
	}
	for _, value := range fileResult {
		for fileName, scheduleTime := range value {
			log.Println(fileName, scheduleTime)
			t, err := time.Parse("2006-01-02 15:04:05", strings.Split(scheduleTime, ".")[0])
			if err != nil {
				log.Panic(err)
			}

			now := time.Now()
			if now.Hour() == t.Hour() && now.Minute() == t.Minute() {
				log.Println("Match found at >>>>>>>>> ", t)
			} else {
				log.Println("Out of match at", now)
			}
		}
	}
	// Update the next execution time and save
	return nil
}
