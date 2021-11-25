package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store"
)

func main() {
	hs := handler.Student{
		Store: store.NewMemoryStudent(),
	}

	app := fiber.New()
	g := app.Group("/")
	hs.Register(g)

	if err := app.Listen(":1373"); err != nil {
		log.Println("cannot start the server")
	}
}
