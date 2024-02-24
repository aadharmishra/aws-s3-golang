package config

import (
	"awsS3Golang/config/models"
	"encoding/json"
	"log"
	"os"
)

var GlobalConfig *models.Config

func InitConfig() error {
	var err error
	var data []byte

	configFile := "/Users/aadharmishra/Documents/github/aws-s3-golang/config.json"
	data, err = os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &GlobalConfig)

	if err != nil || GlobalConfig == nil {
		log.Fatalf("Error getting creds: %s", err)
		return err
	}

	return nil
}
