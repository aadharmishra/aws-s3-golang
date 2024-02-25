package services

import (
	"awsS3Golang/config/models"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
)

type AwsClientInterface interface {
	UploadDoc(ctx *gin.Context) (bool, error)
	GetDoc(ctx *gin.Context) (bool, error)
}

type S3Client struct {
	Config *models.Config
}

func NewS3Client(config *models.Config) AwsClientInterface {
	return &S3Client{
		Config: config,
	}
}

var GlobalAwsSessionPool map[string]*s3.S3

func (client *S3Client) InitS3Client(ctx *gin.Context) (*s3.S3, error) {

	var err error
	if GlobalAwsSessionPool["s3Session"] != nil {
		return GlobalAwsSessionPool["s3Session"], nil
	}

	err = os.Setenv("AWS_ACCESS_KEY_ID", client.Config.Aws.AwsAccessKey)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return nil, err
	}
	err = os.Setenv("AWS_SECRET_ACCESS_KEY", client.Config.Aws.AwsSecretKey)
	if err != nil {
		fmt.Println("Error setting environment variable:", err)
		return nil, err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(client.Config.Aws.AwsRegion),
	})

	if sess == nil || err != nil {
		fmt.Printf("session is empty or error is session creation")
		return nil, err
	}

	svc := s3.New(sess)

	if svc == nil {
		fmt.Printf("empty service object")
		return nil, err
	}

	GlobalAwsSessionPool = make(map[string]*s3.S3)
	GlobalAwsSessionPool["s3Session"] = svc

	return svc, nil
}

func (client *S3Client) UploadDoc(ctx *gin.Context) (bool, error) {

	var err error

	svc, err := client.InitS3Client(ctx)

	if err != nil {
		fmt.Printf("error while s3 client init")
		return false, err
	}

	file, header, err := ctx.Request.FormFile("document")
	if err != nil {
		fmt.Printf("error while extracting request form data")
		return false, err
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, file)
	if err != nil {
		fmt.Printf("error while reading file buffer")
		return false, err
	}

	fileKey := client.Config.Aws.UploadFolderPath + "/" + header.Filename

	input := &s3.PutObjectInput{
		Bucket: aws.String(client.Config.Aws.S3BucketName),
		Key:    aws.String(fileKey),
		Body:   bytes.NewReader(buf.Bytes()),
	}

	_, err = svc.PutObject(input)
	if err != nil {
		fmt.Printf("error while uploading doc to s3")
		return false, err
	}

	return true, nil
}

func (client *S3Client) GetDoc(ctx *gin.Context) (bool, error) {
	var err error

	svc, err := client.InitS3Client(ctx)

	if err != nil {
		fmt.Printf("error while s3 client init")
		return false, err
	}

	fileName := ctx.GetHeader("key")

	docKey := client.Config.Aws.UploadFolderPath + "/" + fileName

	input := &s3.GetObjectInput{
		Bucket: aws.String(client.Config.Aws.S3BucketName),
		Key:    aws.String(docKey),
	}

	videoObject, err := svc.GetObject(input)

	if err != nil || videoObject == nil {
		fmt.Printf("error while getting s3 doc or output is empty")
		return false, err
	}
	defer videoObject.Body.Close()

	videoBuffer := new(bytes.Buffer)
	_, err = videoBuffer.ReadFrom(videoObject.Body)
	if err != nil {
		fmt.Printf("error while reading output from s3")
		return false, err
	}

	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=%s", docKey))
	ctx.Data(http.StatusOK, "video/mp4", videoBuffer.Bytes())

	return false, nil
}
