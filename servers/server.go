package servers

import (
	"awsS3Golang/apis/aws"
	"awsS3Golang/config"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer() error {
	address := config.GlobalConfig.Server.Address
	router := gin.Default()

	corsConfig := cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
	}

	router.Use(cors.New(corsConfig))

	aws.NewAwsRoutes(router, config.GlobalConfig)

	err := router.Run(address)
	if err != nil {
		return err
	}

	fmt.Printf("HTTP Server listening on : " + address)

	return nil
}
