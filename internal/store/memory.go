package store

import (
	"context"

	"githuh.com/cng-by-example/students/internal/model"
)

// MemoryStudent is an implementation of Student store which uses memory as storage.
type MemoryStudent struct {
	students map[string]model.Student
}

func NewMemoryStudent() *MemoryStudent {
	return &MemoryStudent{
		students: make(map[string]model.Student),
	}
}

func (ms *MemoryStudent) Save(_ context.Context, student model.Student) error {
	if _, ok := ms.students[student.ID]; ok {
		return ErrSutdentDuplicate
	}

	ms.students[student.ID] = student

	return nil
}

func (ms *MemoryStudent) LoadByID(_ context.Context, id string) (model.Student, error) {
	s, ok := ms.students[id]
	if !ok {
		return model.Student{}, ErrStudentNotFound
	}

	return s, nil
}

func (ms *MemoryStudent) Load(_ context.Context) ([]model.Student, error) {
	ss := make([]model.Student, 0, len(ms.students))

	for _, s := range ms.students {
		ss = append(ss, s)
	}

	return ss, nil
}
