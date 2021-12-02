package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/db"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store"
)

func main() {
	db, err := db.New(db.Config{
		URL:               "mongodb://127.0.0.1:27017",
		Name:              "sbu",
		ConnectionTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	hs := handler.Student{
		Store: store.NewMongoDBStore(db),
	}

	app := fiber.New()
	g := app.Group("/")
	hs.Register(g)

	if err := app.Listen(":1373"); err != nil {
		log.Println("cannot start the server")
	}
}
