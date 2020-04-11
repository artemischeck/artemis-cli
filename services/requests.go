package services

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

const globalConfigName = "artemis.ini"

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
	start := time.Now()
	var status int
	status, statusText, err := serviceFile.sendAPIRequest()
	duration := time.Since(start)

	// Package response and trigger healthcheck request
	if serviceFile.ServiceType == "REST" {
		statusText = http.StatusText(status)
	}
	log.Println(statusText)
	log.Println(duration)
	// go SendHealthCheck(serviceFile, status, statusText, duration)
}

// SendHealthCheck health check request
func SendHealthCheck(serviceFile ServiceFile, status int, details string, duration time.Duration) {
	fileResult, err := ReadConfigFile(path.Join(ConfigDir, globalConfigName))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fileResult)

	configFile := ConfigFile{}
	configFile.readFile(fileResult)

	// Compose body request
	success := 4
	if status == http.StatusAccepted || status == http.StatusOK || status == http.StatusCreated {
		success = 1
	}
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	hlt := HealthCheckPayload{}
	hlt.Label = serviceFile.Label
	hlt.Status = success
	hlt.Host = host
	hlt.DateTime = time.Now()
	hlt.Message = details
	hlt.Duration = duration

	// Send config data
	var body []byte
	body, err = json.Marshal(hlt)
	if err != nil {
		log.Fatal(err)
	}

	_, _, err = configFile.sendAPIRequest(body)
	if err != nil {
		log.Println(err)
	}
}
