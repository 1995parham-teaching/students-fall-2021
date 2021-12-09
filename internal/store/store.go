package store

import (
	"context"
	"errors"

	"githuh.com/cng-by-example/students/internal/model"
)

var (
	ErrStudentNotFound  = errors.New("student with given id doesn't exist")
	ErrSutdentDuplicate = errors.New("student with given id already exists")
)

type Student interface {
	Save(context.Context, model.Student) error
	LoadByID(context.Context, string) (model.Student, error)
	Load(context.Context) ([]model.Student, error)
}
