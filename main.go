package main

import (
	"log"
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/handlers"
	"todo-app/internal/infrastructure"
	"todo-app/internal/usecases"
	"github.com/gorilla/mux"
)

func main() {
	db := infrastructure.NewMySQLDB(config.DBConfig)
	redisClient := infrastructure.NewRedisClient(config.RedisConfig)
	s3Client := config.NewS3Client(config.S3Config)

	todoRepo := infrastructure.NewMySQLTodoRepository(db)
	fileStorage := infrastructure.NewS3Storage(s3Client, "my-app-files-service")
	redisStream := infrastructure.NewRedisStreamPublisher(redisClient)
	defer redisStream.Close()

	todoService := usecases.NewTodoService(todoRepo, fileStorage, redisStream)
	fileService := usecases.NewFileService(fileStorage)

	todoHandler := handlers.NewTodoHandler(todoService)
	fileHandler := handlers.NewFileHandler(fileService)

	r := mux.NewRouter()
	r.HandleFunc("/todo", todoHandler.CreateTodoItem).Methods("POST")
	r.HandleFunc("/upload", fileHandler.UploadFile).Methods("POST")

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
