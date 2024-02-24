package services

import (
	"awsS3Golang/config/models"
	"bytes"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type AwsClientInterface interface {
	UploadDoc(ctx *gin.Context) (bool, error)
}

type S3Client struct {
	Config *models.Config
}

func NewS3Client(config *models.Config) AwsClientInterface {
	return &S3Client{
		Config: config,
	}
}

func (client *S3Client) UploadDoc(ctx *gin.Context) (bool, error) {

	os.Setenv("AWS_ACCESS_KEY_ID", client.Config.Aws.AwsAccessKey)
	os.Setenv("AWS_SECRET_ACCESS_KEY", client.Config.Aws.AwsSecretKey)

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(client.Config.Aws.AwsRegion),
		Credentials: aws.NewConfig().Credentials,
	})

	if sess == nil || err != nil {
		return false, err
	}

	svc := s3.New(sess)

	file, header, err := ctx.Request.FormFile("document")
	if err != nil {
		return false, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		return false, err
	}

	fileKey := client.Config.Aws.UploadFolderPath + header.Filename

	input := &s3.PutObjectInput{
		Bucket: aws.String(client.Config.Aws.S3BucketName),
		Key:    aws.String(fileKey),
		Body:   bytes.NewReader(buf.Bytes()),
	}

	_, err = svc.PutObject(input)
	if err != nil {
		return false, err
	}

	return true, nil
}
