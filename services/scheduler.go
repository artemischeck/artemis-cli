package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var logDir string = "."

// ScheduleTasks register file details for next schedule in schedule.log file
func ScheduleTasks(files []ServiceFile) error {
	var scheduleItems []map[string]string
	for _, file := range files {
		scheduleItem := make(map[string]string)
		nextExecutionTime := time.Now().Add(time.Second * time.Duration(file.Interval)).String()
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

	file.Truncate(0)
	file.Seek(0, 0)
	_, err = file.WriteString(string(line))
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
}

// RescheduleFiles collect executed files and put them back to the queue
func RescheduleFiles(fileNames []string) {
	var serviceFiles []ServiceFile
	for _, fileName := range fileNames {
		fileResult, err := ReadConfigFile(path.Join(ConfigDir, "conf.d", fileName))
		if err != nil {
			log.Fatal(err)
		}
		serviceFile := ServiceFile{}
		serviceFile.readFile(fileName, fileResult)
		serviceFiles = append(serviceFiles, serviceFile)
	}
	if len(serviceFiles) > 0 {
		ScheduleTasks(serviceFiles)
	}
}
