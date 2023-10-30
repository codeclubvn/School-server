package s3

import (
	"elearning/config"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Service struct {
	BucketName string
	Domain     string
	*s3.S3
}

func ConnectS3(cfg *config.Environment) (*S3Service, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.S3Region),
		Credentials: credentials.NewStaticCredentials(cfg.S3ApiKey, cfg.S3SecretKey, ""),
	})

	if err != nil {
		log.Fatalln(err)
	}

	service := s3.New(session)

	return &S3Service{cfg.S3BucketName, cfg.S3Domain, service}, nil
}
