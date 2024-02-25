package config

import (
	"awsS3Golang/config/models"
	"encoding/json"
	"fmt"
	"os"
)

var GlobalConfig *models.Config

func InitConfig() error {
	var err error
	var data []byte

	configFile := "/Users/aadharmishra/Documents/github/aws-s3-golang/config.json"
	data, err = os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("error while reading config.")
		return err
	}

	err = json.Unmarshal(data, &GlobalConfig)

	if err != nil || GlobalConfig == nil {
		fmt.Printf("error while mapping config.")
		return err
	}

	return nil
}
