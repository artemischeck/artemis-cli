package services

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// ConfigFile schema definition
type ConfigFile struct {
	URL     string
	Secure  bool
	Version string
	Key     string
}

func (file *ConfigFile) readFile(data map[string]string) {
	var err error
	secure := true
	version := "v1"

	if data["URL"] == "" {
		log.Panicln("Validation error, config file: URL is required")
	}
	if data["SECURE"] != "" {
		secure, err = strconv.ParseBool(data["SECURE"])
		if err != nil {
			log.Panicln("Value error, config file", err)
		}
	}
	if data["VERSION"] == "" {
		version = data["VERSION"]
	}

	file.URL = data["URL"]
	file.Secure = secure
	file.Version = version
	file.Key = data["KEY"]
}

func (file *ConfigFile) sendAPIRequest(body []byte) (int, interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", file.URL, bytes.NewBuffer(body))
	req.Header = map[string][]string{"Content-Type": {"application/json"}, "Authorization": {"Token " + file.Key}}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	log.Println(file.URL+" Response: ", resp)
	return resp.StatusCode, resp.Body, err
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

func (file *ServiceFile) sendAPIRequest() (int, interface{}, error) {
	switch serviceType := file.ServiceType; serviceType {
	case "REST":
		if file.Request == "POST" {
			client := &http.Client{}
			body := []byte(file.Data)
			req, err := http.NewRequest("POST", file.URL, bytes.NewBuffer(body))
			if file.AuthType == "BASIC" && file.AuthData != "" {
				authData := strings.Split(file.AuthData, ":")
				req.SetBasicAuth(authData[0], authData[1])
			}
			if file.ContentType != "" {
				req.Header = map[string][]string{"Content-Type": {file.ContentType}}
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}
			log.Println(file.URL+" Response: ", resp)
			return resp.StatusCode, resp.Body, err
		}
		resp, err := http.Get(file.URL)
		if err != nil {
			log.Println(err)
		}
		log.Println(file.URL+" Response: ", resp)
		return resp.StatusCode, resp.Body, err
	case "TELNET":
		log.Panicln("Perform telnet")
		return 0, "", nil
	case "SOAP":
		log.Panicln("Perform SOAP")
		return 0, "", nil
	case "UTIL":
		log.Panicln("Perform UTIL i.e trigger command")
		return 0, "", nil
	default:
		return 0, "", nil
	}
}
