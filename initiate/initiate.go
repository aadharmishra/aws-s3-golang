package initiate

import (
	"awsS3Golang/config"
	"awsS3Golang/servers"
	"fmt"
)

func Initialize() error {
	err := config.InitConfig()
	if err != nil {
		fmt.Printf("config init failed.")
		return err
	}

	// Initializes servers
	err = servers.InitServer()
	if err != nil {
		fmt.Printf("server init failed.")
		return err
	}

	return nil
}
