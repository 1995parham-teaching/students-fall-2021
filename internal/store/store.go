package store

import (
	"errors"

	"githuh.com/cng-by-example/students/internal/model"
)

var (
	ErrStudentNotFound  = errors.New("student with given id doesn't exist")
	ErrSutdentDuplicate = errors.New("student with given id already exists")
)

type Student interface {
	Save(model.Student) error
	LoadByID(id string) (model.Student, error)
	Load() ([]model.Student, error)
}
