package handlers

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	pb "todo"
	"todo/todo"
)

var logger2 log.Logger

// NewService returns a naive, stateless implementation of Service.
func NewService() pb.TodoServer {
	ioWriter := log.New(os.Stdout, "\r\n", log.LstdFlags)
	ioWriter.Println("started user-uas-svc")

	dbHost := os.Getenv("DB_HOST")
	dbUserName := os.Getenv("DB_USER")
	dbSecret := os.Getenv("DB_SECRET")
	dbName := os.Getenv("DB_NAME")

	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=UTC", dbUserName, dbSecret, dbHost, dbName)), &gorm.Config{
		Logger: logger.New(ioWriter,
			logger.Config{
				SlowThreshold: time.Millisecond * 200,
				LogLevel:      0,
			},
		),
	})
	if err != nil {
		println("database is not reachable", "error", err)
		os.Exit(3)
	}

	err = db.AutoMigrate(&todo.Todo{})
	if err != nil {
		println("failed to migrate db", "error", err)
		os.Exit(7)
	}

	/* DOMAIN LOGIC */
	todoRepository := todo.NewTodoRepository(db)
	todoSvc := todo.NewTodoService(todoRepository, &logger2)

	return todoService{
		todoManager: todoSvc,
	}
}

type todoService struct {
	todoManager todo.TodoService
}

func (s todoService) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	resp, err := s.todoManager.GetTodo(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTodoResponse{Todo: resp.ToPb()}, nil
}

func (s todoService) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	succ, err := s.todoManager.DeleteTodo(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTodoResponse{Success: succ}, nil
}

func (s todoService) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	input := todo.Todo{
		User: uint(in.User),
		Todo: in.Todo,
		Due:  in.Due,
		Done: false,
	}

	resp, err := s.todoManager.CreateTodo(ctx, &input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTodoResponse{Todo: resp.ToPb()}, nil
}

func (s todoService) GetAll(ctx context.Context, in *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	resp, err := s.todoManager.GetAllTodos(ctx)
	if err != nil {
		return nil, err
	}

	ret := make([]*pb.Task, len(resp))
	for i, value := range resp {
		ret[i] = value.ToPb()
	}

	return &pb.GetAllResponse{Todo: ret}, nil
}
