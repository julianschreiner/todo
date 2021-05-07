package todo

import (
	"context"
	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	GetAllTodos(ctx context.Context) ([]*Todo, error)
	GetTodoById(ctx context.Context, id uint64) (*Todo, error)
	DeleteTodoBy(ctx context.Context, id uint64) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	}
}

func (r *todoRepository) CreateTodo(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}

func (r *todoRepository) GetAllTodos(ctx context.Context) ([]*Todo, error) {
	return nil, nil
}

func (r *todoRepository) GetTodoById(ctx context.Context, id uint64) (*Todo, error) {
	return nil, nil
}

func (r *todoRepository) DeleteTodoBy(ctx context.Context, id uint64) error {
	return nil
}
