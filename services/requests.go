package services

import (
	"log"
	"os"
	"path"
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
	var globalConfig os.FileInfo
	fileResult, err := ReadConfigFile(path.Join(ConfigDir, globalConfig.Name()))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fileResult)

	configFile := ConfigFile{}
	configFile.readFile(fileResult)

	// Send config data
	_, _, err = configFile.sendAPIRequest([]byte("s"))
	if err != nil {
		log.Fatal(err)
	}
}
