package interfaces

import (
	"context"
	"todo-app/internal/entities"
)

type RedisStream interface {
	Publish(ctx context.Context, item *entities.TodoItem) error
	Subscribe(ctx context.Context) (<-chan *entities.TodoItem, error)
	Close() error
}
