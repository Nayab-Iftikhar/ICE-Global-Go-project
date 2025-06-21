package usecases

import (
	"context"
	"time"
	"todo-app/internal/entities"
	"todo-app/internal/interfaces"
	"github.com/google/uuid"
)

type TodoService struct {
	todoRepo    interfaces.TodoRepository
	fileStorage interfaces.FileStorage
	redisStream interfaces.RedisStream
}

func NewTodoService(todoRepo interfaces.TodoRepository, fileStorage interfaces.FileStorage, redisStream interfaces.RedisStream) *TodoService {
	return &TodoService{
		todoRepo:    todoRepo,
		fileStorage: fileStorage,
		redisStream: redisStream,
	}
}

func (s *TodoService) CreateTodoItem(ctx context.Context, description string, dueDate time.Time, fileID string) (*entities.TodoItem, error) {
	todoItem := &entities.TodoItem{
		ID:          generateUUID(),
		Description: description,
		DueDate:     dueDate,
		FileID:      fileID,
	}

	if err := s.todoRepo.Save(ctx, todoItem); err != nil {
		return nil, err
	}

	if err := s.redisStream.Publish(ctx, todoItem); err != nil {
		return nil, err
	}

	return todoItem, nil
}

func generateUUID() string {
	return uuid.New().String()
}
