package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"structs/shared"
)

func GetConfiguration() shared.Configuration{
	var configuration shared.Configuration
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Unable to open config.json !")
	}

	fileContent, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(fileContent, &configuration)
	_ = file.Close()

	return configuration
}