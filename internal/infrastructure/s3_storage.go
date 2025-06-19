package infrastructure

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	client     *s3.S3
	bucketName string
}

func NewS3Storage(client *s3.S3, bucketName string) *S3Storage {
	
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil && !isBucketExistsError(err) {
		log.Fatalf("Failed to create bucket: %v", err)
	}
	return &S3Storage{
		client:     client,
		bucketName: bucketName,
	}
}


func isBucketExistsError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "BucketAlreadyOwnedByYou") ||
		strings.Contains(err.Error(), "BucketAlreadyExists")
}

func (s *S3Storage) UploadFile(ctx context.Context, file []byte) (string, error) {
	filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), "file")

	input := &s3.PutObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	}

	_, err := s.client.PutObjectWithContext(ctx, input)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		log.Fatal("Failed to create AWS session:", err)
	}

	s3Client := s3.New(sess)
	storage := NewS3Storage(s3Client, "my-app-files-service")

	
	ctx := context.Background()
	fileBytes := []byte("hello world from Go")
	key, err := storage.UploadFile(ctx, fileBytes)
	if err != nil {
		log.Fatal("Upload error:", err)
	}

	fmt.Println("Uploaded successfully with key:", key)
}
