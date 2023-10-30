package repository

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"

	service "elearning/infra/s3"
)

const (
	filePreSignExpireDuration = time.Hour * 12
)

type s3Repository struct {
	service *service.S3Service
}

func NewS3Repository(service *service.S3Service) *s3Repository {
	return &s3Repository{
		service: service,
	}
}

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"column:cloud_name;"`
}

func (repo *s3Repository) SaveImageUploaded(ctx context.Context, data []byte) (*Image, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	uploadID := uuid.New().String()

	_, err := repo.service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(repo.service.BucketName),
		Key:         aws.String(uploadID),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return nil, err
	}

	img := &Image{
		Url:       fmt.Sprintf("%s/%s", repo.service.Domain, uploadID),
		CloudName: "s3",
	}

	return img, err
}

func (repo *s3Repository) SaveFile(ctx context.Context, data []byte, dst string) (string, error) {
	fileBytes := bytes.NewReader(data)
	fileType := http.DetectContentType(data)

	uploadID := uuid.New().String()

	_, err := repo.service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(repo.service.BucketName),
		Key:         aws.String(uploadID),
		ACL:         aws.String("private"),
		ContentType: aws.String(fileType),
		Body:        fileBytes,
	})

	if err != nil {
		return "", err
	}

	return uploadID, nil
}

func (repo *s3Repository) GetFileUrl(ctx context.Context, uploadID string) (string, error) {

	req, _ := repo.service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(repo.service.BucketName),
		Key:    aws.String(uploadID),
	})

	url, err := req.Presign(filePreSignExpireDuration)

	if err != nil {
		return "", err
	}

	return url, nil
}
