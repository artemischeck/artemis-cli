package services

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const globalConfigName = "healthcheck.ini"

//Register opens services file and run based on settings
func Register(dir string) error {
	var globalConfig os.FileInfo
	var fileResult map[string]string
	var serviceFiles []ServiceFile
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if file.Name() == globalConfigName {
			globalConfig = file
			continue
		}
		// Read service files
		fileName := file.Name()
		fileResult, err = readConfigFile(path.Join(dir, fileName))
		serviceFile := ServiceFile{}
		serviceFile.readFile(fileName, fileResult)
		serviceFiles = append(serviceFiles, serviceFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Read global config file
	// log.Println("Config", globalConfig.Name())
	fileResult, err = readConfigFile(path.Join(dir, globalConfig.Name()))

	// Schedule tasks
	if len(serviceFiles) > 0 {
		ScheduleTasks(serviceFiles)
	}
	return nil
}

func readConfigFile(fileName string) (map[string]string, error) {
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
