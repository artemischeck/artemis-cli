package services

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var (
	// ConfigDir sent config files directory
	ConfigDir string
)

//Register opens services file and run based on settings
func Register() error {
	var fileResult map[string]string
	var serviceFiles []ServiceFile
	files, err := ioutil.ReadDir(ConfigDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == globalConfigName {
			continue
		}
		// Read service files
		fileName := file.Name()
		fileResult, err = ReadConfigFile(path.Join(ConfigDir, fileName))
		if err != nil {
			log.Fatal(err)
		}
		serviceFile := ServiceFile{}
		serviceFile.readFile(fileName, fileResult)
		serviceFiles = append(serviceFiles, serviceFile)

	}

	// Schedule tasks
	if len(serviceFiles) > 0 {
		ScheduleTasks(serviceFiles)
	}
	return nil
}

// ReadConfigFile config files
func ReadConfigFile(fileName string) (map[string]string, error) {
	results := make(map[string]string)
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip new lines
		if line == "" {
			continue
		}
		// Skip comments
		if strings.HasPrefix(line, "#") {
			continue
		}
		config := strings.Split(line, " ")
		results[config[0]] = config[1]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
