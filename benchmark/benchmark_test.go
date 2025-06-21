package benchmark

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func BenchmarkMySQLInsert(b *testing.B) {

	dsn := "root:password@tcp(localhost:3306)/todoapp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	db.Exec("TRUNCATE TABLE todo_items")

	for i := 0; i < b.N; i++ {
		_, err := db.Exec(`INSERT INTO todo_items (id, description, due_date, file_id) VALUES (?, ?, ?, ?)`,
			fmt.Sprintf("bench-%d", i), fmt.Sprintf("desc %d", i), time.Now(), "file-id")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkS3Upload(b *testing.B) {
	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2"),
		Credentials: credentials.NewStaticCredentials(
			"test", "test", ""), // from localstack env
		Endpoint:         aws.String("http://localhost:4566"),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		b.Fatal(err)
	}

	s3Client := s3.New(awsSession)
	data := []byte("benchmark file content")

	_, _ = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String("my-app-files-service"),
	})

	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("benchmarks/file-%d.txt", i)
		_, err := s3Client.PutObject(&s3.PutObjectInput{
			Bucket: aws.String("my-app-files-service"),
			Key:    aws.String(key),
			Body:   bytes.NewReader(data),
		})
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRedisStreamPublish(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	ctx := context.Background()

	for i := 0; i < b.N; i++ {
		_, err := rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: "mystream",
			Values: map[string]interface{}{"message": fmt.Sprintf("event %d", i)},
		}).Result()
		if err != nil {
			b.Fatal(err)
		}
	}
}
