package models

type Config struct {
	Server Server `json:"server,omitempty"`
	Aws    Aws    `json:"aws,omitempty"`
}

type Server struct {
	Address string `json:"address,omitempty"`
}

type Aws struct {
	AwsRegion        string `json:"awsRegion"`
	AwsAccessKey     string `json:"awsAccessKey"`
	AwsSecretKey     string `json:"awsSecretKey"`
	S3BucketName     string `json:"s3BucketName"`
	S3ObjectPrefix   string `json:"s3ObjectPrefix"`
	ServerPort       string `json:"serverPort"`
	UploadFolderPath string `json:"uploadFolderPath"`
}
