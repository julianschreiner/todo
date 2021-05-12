package handlers

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger "gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
	pb "todo"
	"todo/auth"
	"todo/todo"
	"todo/user"
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
	userClient := user.NewUserClient()

	todoRepository := todo.NewTodoRepository(db)
	todoSvc := todo.NewTodoService(todoRepository, &logger2)

	authClient := auth.NewauthClient()

	return todoService{
		todoManager: todoSvc,
		authClient:  authClient,
		userClient:  userClient,
	}
}

type todoService struct {
	todoManager todo.TodoService
	authClient  auth.AuthClient
	userClient  user.UserClient
}

func (s todoService) validateCookie(ctx context.Context) string {
	rawCookies := ctx.Value("cookie")
	token := ""

	if rawCookies != nil {
		header := http.Header{}
		header.Add("Cookie", rawCookies.(string))
		request := http.Request{Header: header}
		c, _ := request.Cookie("refresh")
		if c != nil {
			token = c.Value
		}
	}

	return token
}

func (s todoService) GetTodo(ctx context.Context, in *pb.GetTodoRequest) (*pb.GetTodoResponse, error) {
	// permission check
	token := s.validateCookie(ctx)

	err := s.authClient.Validate(token)
	if err != nil {
		// not authorized
		return nil, errors.New(err.Error())
	}

	resp, err := s.todoManager.GetTodo(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTodoResponse{Todo: resp.ToPb()}, nil
}

func (s todoService) DeleteTodo(ctx context.Context, in *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	// permission check
	token := s.validateCookie(ctx)

	err := s.authClient.Validate(token)
	if err != nil {
		// not authorized
		return nil, errors.New(err.Error())
	}

	succ, err := s.todoManager.DeleteTodo(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTodoResponse{Success: succ}, nil
}

func (s todoService) CreateTodo(ctx context.Context, in *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	// permission check
	token := s.validateCookie(ctx)

	err := s.authClient.Validate(token)
	if err != nil {
		// not authorized
		return nil, errors.New(err.Error())
	}

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
	// permission check
	token := s.validateCookie(ctx)

	err := s.authClient.Validate(token)
	if err != nil {
		// not authorized
		return nil, errors.New(err.Error())
	}

	resp, err := s.todoManager.GetAllTodosForUser(ctx, in.User)
	if err != nil {
		return nil, err
	}

	ret := make([]*pb.Task, len(resp))
	for i, value := range resp {
		ret[i] = value.ToPb()
	}

	/* FETCH USER INFORMATION FROM USER SERVICE */
	var usr *user.User
	if len(ret) > 0 {
		usr, err = s.userClient.GetActiveUserById(in.User)
		if err != nil {
			return nil, errors.New(err.Error())
		}

	}

	return &pb.GetAllResponse{User: usr.ToPb(), Todo: ret}, nil
}

func (s todoService) UpdateTodo(ctx context.Context, in *pb.UpdateTodoRequest) (*pb.UpdateTodoResponse, error) {
	var resp pb.UpdateTodoResponse
	return &resp, nil
}
