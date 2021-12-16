package config

import (
	"time"

	"githuh.com/cng-by-example/students/internal/db"
)

func Default() Config {
	return Config{
		Database: db.Config{
			URL:               "mongodb://127.0.0.1:27017",
			ConnectionTimeout: 10 * time.Second,
			Name:              "sbu",
		},
	}
}
