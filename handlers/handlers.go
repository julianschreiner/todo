package handlers

import (
	"context"

	pb "todo"
)

// NewService returns a naive, stateless implementation of Service.
func NewService() pb.TodoServer {
	return todoService{}
}

type todoService struct{}

func (s todoService) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	var resp pb.CreateResponse
	return &resp, nil
}
