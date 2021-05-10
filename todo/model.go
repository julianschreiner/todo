package todo

import (
	"gorm.io/gorm"
	pb "todo"
)

type Todo struct {
	gorm.Model
	User uint
	Todo string
	Due  string
	Done bool
}

func (t *Todo) ToPb() *pb.Task {
	return &pb.Task{
		User: uint64(t.User),
		Todo: t.Todo,
		Due:  t.Due,
		Done: t.Done,
	}
}
