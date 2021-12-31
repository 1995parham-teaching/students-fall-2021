package main

import (
	"context"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/config"
	"githuh.com/cng-by-example/students/internal/db"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("you must specify a mode")

		return
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Println("cannot create a logger")

		return
	}

	cfg := config.New()

	switch os.Args[1] {
	case "server":
		server(cfg, logger.Named("server"))
	case "migrate":
		migrate(cfg, logger.Named("migrate"))
	default:
		log.Println("you must specify a mode between server or migrate")
	}
}

func server(cfg config.Config, logger *zap.Logger) {
	db, err := db.New(cfg.Database)
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	logger.Info("database connection successful")

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

func migrate(cfg config.Config, logger *zap.Logger) {
	db, err := db.New(cfg.Database)
	if err != nil {
		log.Fatalf("database connection failed %s", err)
	}

	logger.Info("database connection successful")

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
