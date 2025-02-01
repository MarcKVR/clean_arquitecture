package s3

import (
	"bytes"
	"mime/multipart"

	"github.com/MarcKVR/clean_arquitecture/cmd/config"
	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3ClientImp struct {
	client     *s3.S3
	bucketName string
}

func NewS3SClient(cfg *config.Config) repositories.FileRepository {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, ""),
	}))

	return &S3ClientImp{
		client:     s3.New(sess),
		bucketName: cfg.S3BucketName,
	}
}

func (s *S3ClientImp) UploadFile(file *multipart.FileHeader, destination string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(src)
	if err != nil {
		return "", err
	}

	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(destination),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	url := "https://" + s.bucketName + ".s3.amazonaws.com/" + destination

	return url, nil
}
