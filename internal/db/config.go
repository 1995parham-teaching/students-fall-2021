package db

import "time"

type Config struct {
	URL               string
	ConnectionTimeout time.Duration
	Name              string
}
