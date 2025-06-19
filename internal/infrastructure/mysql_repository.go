package infrastructure

import (
	"context"
	"todo-app/internal/entities"

	"gorm.io/gorm"
)

type MySQLTodoRepository struct {
	db *gorm.DB
}

func NewMySQLTodoRepository(db *gorm.DB) *MySQLTodoRepository {
	return &MySQLTodoRepository{db: db}
}

func (r *MySQLTodoRepository) Save(ctx context.Context, item *entities.TodoItem) error {
	return r.db.Create(item).Error
}

func (r *MySQLTodoRepository) GetByID(ctx context.Context, id string) (*entities.TodoItem, error) {
	var item entities.TodoItem
	err := r.db.First(&item, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}
