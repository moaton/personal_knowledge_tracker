package dto

import (
	"time"
)

type Resource struct {
	ID        string
	Title     string
	Type      string
	Content   string
	Tags      []string
	Metadata  map[string]interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}
