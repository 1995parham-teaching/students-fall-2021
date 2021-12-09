package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/db"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("you must specify a mode")

		return
	}

	switch os.Args[1] {
	case "server":
		server()
	case "migrate":
		migrate()
	default:
		log.Println("you must specify a mode between server or migrate")
	}
}

func server() {
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

func migrate() {
	db, err := db.New(db.Config{
		URL:               "mongodb://127.0.0.1:27017",
		Name:              "sbu",
		ConnectionTimeout: 10 * time.Second,
	})
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	name, err := db.Collection(store.Collection).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.M{"id": 1},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatalf("cannot create an index %s", err)
	}

	log.Println(name)
}
