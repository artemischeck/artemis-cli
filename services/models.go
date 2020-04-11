package services

import (
	"bytes"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// HealthCheckPayload request schema
type HealthCheckPayload struct {
	Label    string        `json:"label"`
	Status   int           `json:"status"`
	DateTime time.Time     `json:"date_time"`
	Duration time.Duration `json:"duration"`
	Message  string        `json:"message"`
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

func (file *ConfigFile) sendAPIRequest(body []byte) (int, string, error) {
	client := &http.Client{Timeout: time.Duration(10)}
	req, err := http.NewRequest("POST", file.URL, bytes.NewBuffer(body))
	if err != nil {
		return http.StatusServiceUnavailable, "Could not send request", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Token "+file.Key)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	log.Println(file.URL+" Response: ", resp)
	return resp.StatusCode, "", err
}

// ServiceFile schema
type ServiceFile struct {
	Name            string
	Label           string
	ServiceType     string
	AuthType        string
	AuthData        string
	ContentType     string
	URL             string
	Host            string
	Port            int
	Request         string
	Data            string
	Interval        int
	Timeout         int
	CMD             string
	UtilServiceName string
	OSServiceName   string
	MinThreshold    int
	MaxThreshold    int
	Tags            string
}

func (file *ServiceFile) readFile(fileName string, data map[string]string) {
	var err error
	port := 80
	interval := 60
	timeout := 10
	minThreshold := 0
	maxThreshold := 0
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
	if data["MAX_THRESHOLD"] != "" {
		maxThreshold, err = strconv.Atoi(data["MAX_THRESHOLD"])
		if err != nil {
			log.Panicln("Value error for file:"+fileName, err)
		}
	}
	if data["MIN_THRESHOLD"] != "" {
		minThreshold, err = strconv.Atoi(data["MIN_THRESHOLD"])
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
	file.Host = data["HOST"]
	file.Port = port
	file.Request = data["REQUEST"]
	file.Data = data["DATA"]
	file.Interval = interval
	file.Timeout = timeout
	file.CMD = data["CMD"]
	file.UtilServiceName = data["UTIL_SERVICE_NAME"]
	file.OSServiceName = data["OS_SERVICE_NAME"]
	file.MaxThreshold = maxThreshold
	file.MinThreshold = minThreshold
	file.Tags = tags
}

func (file *ServiceFile) sendAPIRequest() (int, string, error) {
	switch serviceType := file.ServiceType; serviceType {
	case "REST":
		tr := &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    5 * time.Second,
			DisableCompression: true,
		}
		if file.Request == "POST" {
			client := &http.Client{Transport: tr, Timeout: time.Duration(file.Timeout)}
			body := []byte(file.Data)
			req, err := http.NewRequest("POST", file.URL, bytes.NewBuffer(body))
			if file.AuthType == "BASIC" && file.AuthData != "" {
				authData := strings.Split(file.AuthData, ":")
				req.SetBasicAuth(authData[0], authData[1])
			}
			if file.ContentType != "" {
				req.Header.Add("Content-Type", file.ContentType)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
			}
			defer resp.Body.Close()
			log.Println(file.URL+" Response: ", resp)
			return resp.StatusCode, "", err
		}
		resp, err := http.Get(file.URL)
		if err != nil {
			log.Println(err)
		}
		log.Println(file.URL+" Response: ", resp)
		return resp.StatusCode, "", err
	case "TELNET":
		log.Println("Perform telnet")
		dialer := net.Dialer{Timeout: time.Duration(file.Timeout)}
		address := file.Host + ":" + strconv.Itoa(file.Port)
		conn, err := dialer.Dial("tcp", address)
		if err != nil {
			log.Println(err)
			return 0, "Failed to start:", err
		}
		defer conn.Close()
		return http.StatusOK, "Running", nil
	case "PLUGIN":
		log.Println("Perform PLUGIN i.e trigger command")
		pluginBashFile, err := filepath.Abs("../scripts/plugin.sh")
		if err != nil {
			log.Println(err)
			return 0, "Could not load plugin script", err
		}
		log.Println("pluginBashFile", pluginBashFile)
		cmd := exec.Command("/bin/sh", pluginBashFile)

		cmd.Env = append(os.Environ(),
			"SERVICE_NAME="+file.UtilServiceName,
		)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
			return 0, "Failed to start", err
		}
		if string(out) != "" {
			return http.StatusServiceUnavailable, string(out), nil
		}
		return http.StatusOK, "Running", nil
	default:
		return 0, "", nil
	}
}
