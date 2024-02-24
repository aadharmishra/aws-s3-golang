package s3

import (
	awsClient "awsS3Golang/clients"
	"awsS3Golang/config/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type S3Service struct {
	AwsClient awsClient.AwsClientInterface
	Config    *models.Config
}

func NewS3Service(config *models.Config) *S3Service {
	return &S3Service{
		AwsClient: awsClient.NewS3Client(config),
		Config:    config,
	}
}

func (s3 *S3Service) UploadDocumentToS3(ctx *gin.Context) error {

	var success bool
	var err error

	err = s3.ValidateUploadRequest(ctx)

	if err != nil {
		return err
	}

	success, err = s3.AwsClient.UploadDoc(ctx)

	if !success || err != nil {
		return err
	}

	return nil
}

func (s3 *S3Service) ValidateUploadRequest(ctx *gin.Context) error {

	if s3.Config == nil {
		return fmt.Errorf("empty config")
	}

	bucketName := s3.Config.Aws.S3BucketName
	region := s3.Config.Aws.AwsRegion

	if bucketName == "" || region == "" {
		return fmt.Errorf("empty bucket name and region")
	}

	return nil
}
