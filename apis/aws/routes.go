package aws

import (
	"awsS3Golang/config/models"

	"github.com/gin-gonic/gin"
)

func NewAwsRoutes(router *gin.Engine, config *models.Config) {
	BindRoutes(router, config)
}

func BindRoutes(router *gin.Engine, config *models.Config) {
	service := NewAwsService(config)
	routerApi := router.Group("/aws/s3")
	{
		routerApi.POST("/upload", service.PostDocumentToS3)
		routerApi.GET("/get", service.GetDocumentFromS3)
	}
}
