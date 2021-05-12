package todo

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *Todo) error
	GetAllTodos(ctx context.Context, user uint64) ([]*Todo, error)
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

func (r *todoRepository) CreateTodo(ctx context.Context, todo *Todo) error {
	return r.db.Save(todo).Error
}

func (r *todoRepository) GetAllTodos(ctx context.Context, user uint64) ([]*Todo, error) {
	var todos []*Todo
	err := r.db.Where("User = ?", user).Find(&todos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("InternalServerError")
	}

	return todos, nil
}

func (r *todoRepository) GetTodoById(ctx context.Context, id uint64) (*Todo, error) {
	var todo *Todo
	err := r.db.First(todo, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("InternalServerError")
	}

	return todo, nil
}

func (r *todoRepository) DeleteTodoBy(ctx context.Context, id uint64) error {
	todo := &Todo{
		Model: gorm.Model{
			ID: uint(id),
		},
	}

	return r.db.Delete(todo).Error
}
