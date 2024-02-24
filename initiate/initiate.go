package initiate

import (
	"awsS3Golang/config"
	"awsS3Golang/servers"
	"log"
)

func Initialize() error {
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("config init failed.")
		return err
	}

	// Initializes servers
	err = servers.InitServer()
	if err != nil {
		return err
	}

	return nil
}
