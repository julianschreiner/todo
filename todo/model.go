package todo

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	User uint
	Todo string
	Due  string
	Done bool
}
