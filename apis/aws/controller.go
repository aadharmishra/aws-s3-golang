package aws

import (
	"awsS3Golang/config/models"
	"awsS3Golang/services/s3"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AwsService struct {
	S3Service *s3.S3Service
}

func NewAwsService(config *models.Config) *AwsService {
	return &AwsService{
		S3Service: s3.NewS3Service(config),
	}
}

func (aws *AwsService) PostDocumentToS3(ctx *gin.Context) {
	err := aws.S3Service.UploadDocumentToS3(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}

func (aws *AwsService) GetDocumentFromS3(ctx *gin.Context) {
	err := aws.S3Service.FetchDocumentFromS3(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
