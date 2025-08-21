package dto

import (
	"time"
)

type Resource struct {
	ID         string
	Title      string
	Type       string
	Content    string
	Tags       []string
	Metadata   []byte
	Created_at time.Time
	Updated_at time.Time
}
