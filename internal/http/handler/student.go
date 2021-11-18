package handler

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/store"
)

type Student struct {
	Store store.Student
}

func (s Student) List(c *fiber.Ctx) error {
	ss, err := s.Store.Load()
	if err != nil {
		log.Printf("cannot load students %s", err)

		return fiber.ErrInternalServerError
	}

	// nolint: wrapcheck
	return c.Status(http.StatusOK).JSON(ss)
}

func (s Student) Register(g fiber.Router) {
	g.Get("/student", s.List)
}
