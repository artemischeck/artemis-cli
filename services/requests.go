package services

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"time"
)

const globalConfigName = "healthcheck.ini"

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
	var status int
	var details interface{}
	status, details, err = serviceFile.sendAPIRequest()

	// Package response and trigger healthcheck request
	go SendHealthCheck(fileName, status, details)
}

// SendHealthCheck health check request
func SendHealthCheck(fileName string, status int, details interface{}) {
	log.Println("Sending SendHealthCheck")
	// Read global config file
	fileResult, err := ReadConfigFile(path.Join(ConfigDir, globalConfigName))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fileResult)

	configFile := ConfigFile{}
	configFile.readFile(fileResult)

	// Compose body request
	success := true
	if status != 200 && status != 201 {
		success = false
	}
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hlt := HealthCheck{}
	hlt.Service = fileName
	hlt.Status = success
	hlt.Host = host
	hlt.DateTime = time.Now()

	// Send config data
	var body []byte
	body, err = json.Marshal(hlt)
	if err != nil {
		log.Fatal(err)
	}
	
	_, _, err = configFile.sendAPIRequest(body)
	if err != nil {
		log.Fatal(err)
	}
}
