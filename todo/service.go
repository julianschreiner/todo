package todo

import (
	"context"
	"log"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *Todo) (*Todo, error)
	GetAllTodos(ctx context.Context) ([]*Todo, error)
	GetTodo(ctx context.Context, id uint64) (*Todo, error)
	DeleteTodo(ctx context.Context, id uint64) (bool, error)
}

type todoService struct {
	todoRepository TodoRepository
	logger         *log.Logger
}

func NewTodoService(repository TodoRepository, log *log.Logger) TodoService {
	return &todoService{
		todoRepository: repository,
		logger:         log,
	}
}

func (s *todoService) CreateTodo(ctx context.Context, todo *Todo) (*Todo, error) {
	return nil, nil
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]*Todo, error) {
	return nil, nil
}

func (s *todoService) GetTodo(ctx context.Context, id uint64) (*Todo, error) {
	return nil, nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id uint64) (bool, error) {
	return false, nil
}
