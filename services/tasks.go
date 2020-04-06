package services

import (
	"bufio"
	"fmt"
	"os"
)

// ExecuteQueue read queue file and execute based on defination
func ExecuteQueue() error {
	file, err := os.Open("schedule.log")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Schedule items:", line)
	}
	return nil
}
