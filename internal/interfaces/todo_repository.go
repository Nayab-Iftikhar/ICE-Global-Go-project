package interfaces

import (
	"context"
	"todo-app/internal/entities"
)

type TodoRepository interface {
	Save(ctx context.Context, item *entities.TodoItem) error
}
