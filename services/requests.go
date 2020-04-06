package services

import (
	"log"
	"path"
)

// SendRequest per file defination
func SendRequest(fileName string) {
	// 1. Readfile
	fileResult, err := ReadConfigFile(path.Join(ConfigDir, fileName))
	if err != nil {
		log.Fatal(err)
	}
	serviceFile := ServiceFile{}
	serviceFile.readFile(fileName, fileResult)
	// 2. Trigger API
	var status bool
	var details string
	status, details, err = serviceFile.sendAPIRequest()

	// Package response and trigger healthcheck request
	go SendHealthCheck(fileName, status, details)
}

// SendHealthCheck health check request
func SendHealthCheck(fileName string, status bool, details string) {
	log.Println("Sending SendHealthCheck")
}
