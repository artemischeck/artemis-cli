package services

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"time"
)

// HealthCheck request schema
type HealthCheck struct {
	Service  string        `json:"service"`
	Status   bool          `json:"status"`
	DateTime time.Time     `json:"date_time"`
	Duration time.Duration `json:"duration"`
	Details  string        `json:"details"`
	Host     string        `json:"host"`
	Tags     string        `json:"tags"`
}

func (hlt *HealthCheck) compose(results map[string]string) error {
	return nil
}

// ServiceFile schema
type ServiceFile struct {
	Name        string
	Label       string
	ServiceType string
	AuthType    string
	AuthData    string
	ContentType string
	URL         string
	Port        int
	Request     string
	Data        string
	Interval    int
	Timeout     int
	CMD         string
	Tags        string
}

func (file *ServiceFile) readFile(fileName string, data map[string]string) {
	var err error
	port := 80
	interval := 60
	timeout := 10
	tags := "default"

	if data["LABEL"] == "" {
		log.Panicln("Validation error for file:"+fileName, "LABEL is required")
	}
	if data["SERVICE_TYPE"] == "" {
		log.Panicln("Validation error for file:"+fileName, "SERVICE_TYPE is required")
	}
	if data["TAGS"] == "" {
		tags = "default"
	}
	if data["PORT"] != "" {
		port, err = strconv.Atoi(data["PORT"])
		if err != nil {
			log.Panicln("Value error for file:"+fileName, err)
		}
	}
	if data["INTERVAL"] != "" {
		interval, err = strconv.Atoi(data["INTERVAL"])
		if err != nil {
			log.Panicln("Value error for file:"+fileName, err)
		}
	}
	// Ensure interval is more than 60 seconds
	if interval < 60 {
		log.Panicln("Value error for file:"+fileName, "Interval must be more than 60 seconds")
	}
	if data["TIMEOUT"] != "" {
		timeout, err = strconv.Atoi(data["TIMEOUT"])
		if err != nil {
			log.Panicln("Value error for file:"+fileName, err)
		}
	}
	if timeout > 300 {
		log.Panicln("Value error for file:"+fileName, "Timeout value cannot be more than 300 seconds")
	}

	file.Name = fileName
	file.Label = data["LABEL"]
	file.ServiceType = data["SERVICE_TYPE"]
	file.AuthType = data["AUTH_TYPE"]
	file.AuthData = data["AUTH_DATA"]
	file.ContentType = data["CONTENT_TYPE"]
	file.URL = data["URL"]
	file.Port = port
	file.Request = data["REQUEST"]
	file.Data = data["DATA"]
	file.Interval = interval
	file.Timeout = timeout
	file.CMD = data["CMD"]
	file.Tags = tags
}

func (file *ServiceFile) sendAPIRequest() (bool, string, error) {
	switch serviceType := file.ServiceType; serviceType {
	case "REST":
		if file.Request == "POST" {
			body := []byte(file.Data)
			resp, err := http.Post(file.URL, file.ContentType, bytes.NewBuffer(body))
			if err != nil {
				log.Println(err)
			}
			log.Println(file.URL+" Response: ", resp)
			return true, "resp", err
		}
		resp, err := http.Get(file.URL)
		if err != nil {
			log.Println(err)
		}
		log.Println(file.URL+" Response: ", resp)
		return true, "resp", err
	case "TELNET":
		log.Panicln("Perform telnet")
		return false, "", nil
	case "SOAP":
		log.Panicln("Perform SOAP")
		return false, "", nil
	case "UTIL":
		log.Panicln("Perform UTIL i.e trigger command")
		return false, "", nil
	default:
		return false, "", nil
	}
}
