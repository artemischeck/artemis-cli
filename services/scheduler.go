package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var logDir string = "."

// ScheduleTasks register file details for next schedule in schedule.log file
func ScheduleTasks(files []ServiceFile) error {
	var scheduleItems []map[string]string
	for _, file := range files {
		scheduleItem := make(map[string]string)
		nextExecutionTime := time.Now().Local().Add(time.Second * time.Duration(60)).String()
		scheduleItem[file.Name] = nextExecutionTime
		scheduleItems = append(scheduleItems, scheduleItem)
	}
	if len(scheduleItems) > 0 {
		scheduleItemsStr, err := json.Marshal(scheduleItems)
		if err != nil {
			return err
		}
		writeToFile(scheduleItemsStr)
	}
	return nil
}

func writeToFile(line []byte) {
	file, err := os.OpenFile("schedule.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.Write(line)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}
