package handler

import (
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/http/request"
	"githuh.com/cng-by-example/students/internal/model"
	"githuh.com/cng-by-example/students/internal/store"
)

type Student struct {
	Store store.Student
}

// nolint: wrapcheck
func (s Student) List(c *fiber.Ctx) error {
	ss, err := s.Store.Load(c.Context())
	if err != nil {
		log.Printf("cannot load students %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(ss)
}

// nolint: wrapcheck
func (s Student) Create(c *fiber.Ctx) error {
	req := new(request.Student)

	if err := c.BodyParser(req); err != nil {
		log.Printf("cannot load student data %s", err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Printf("cannot validate student data %s", err)

		return fiber.ErrBadRequest
	}

	student := model.Student{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Units:       0,
		PassedUnits: 0,
		Average:     0.0,
	}

	if err := s.Store.Save(c.Context(), student); err != nil {
		if errors.Is(err, store.ErrSutdentDuplicate) {
			return fiber.NewError(http.StatusBadRequest, "student already exists")
		}

		log.Printf("cannot save student %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(student)
}

// nolint: wrapcheck
func (s Student) Get(c *fiber.Ctx) error {
	id := c.Params("id", "")
	if err := validation.Validate(id,
		validation.Required,
		is.Digit,
		validation.Length(request.SBUStudentIDLen, request.SBUStudentIDLen),
	); err != nil {
		log.Printf("cannot validate student id %s", err)

		return fiber.ErrBadRequest
	}

	student, err := s.Store.LoadByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, store.ErrStudentNotFound) {
			return fiber.ErrNotFound
		}

		log.Printf("cannot load student %s", err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(student)
}

func (s Student) Register(g fiber.Router) {
	g.Get("/student", s.List)
	g.Get("/student/:id", s.Get)
	g.Post("/student", s.Create)
}
