package db_test

import (
	"testing"
	"time"

	"githuh.com/cng-by-example/students/internal/config"
	"githuh.com/cng-by-example/students/internal/db"
)

func TestDBWithInvalidConfiguration(t *testing.T) {
	t.Parallel()

	db, err := db.New(db.Config{
		URL:               "",
		Name:              "",
		ConnectionTimeout: 1 * time.Second,
	})
	if err == nil {
		t.Error("err should not be nil")
	}

	if db != nil {
		t.Error("db should be nil")
	}
}

func TestDBWithValidConfiguration(t *testing.T) {
	t.Parallel()

	cfg := config.New()

	db, err := db.New(cfg.Database)
	if err != nil {
		t.Error("err should be nil")
	}

	if db == nil {
		t.Error("db should not be nil")
	}
}
