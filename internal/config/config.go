package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var DBConfig = &gorm.Config{}

var RedisConfig = &redis.Options{
	Addr: "redis:6379",
}

var S3Config = &aws.Config{
	Region:           aws.String("ap-southeast-2"),
	Credentials:      credentials.NewEnvCredentials(),
	Endpoint:         aws.String("http://localstack:4566"),
	S3ForcePathStyle: aws.Bool(true),
}
func NewS3Client(cfg *aws.Config) *s3.S3 {
	awsSession := session.Must(session.NewSession(cfg))
	return s3.New(awsSession)
}
