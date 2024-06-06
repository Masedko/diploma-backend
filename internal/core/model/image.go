package model

import (
	"time"

	"github.com/google/uuid"
)

type Image struct {
	ID                uuid.UUID `json:"id"`
	DestroyedObjectID uuid.UUID `json:"destroyed_object_id"`
	FileName          string    `json:"file_name"`
	Path              string    `json:"path"`
	Lat               float64   `json:"lat"`
	Lng               float64   `json:"lng"`
	X                 int       `json:"x"`
	Y                 int       `json:"y"`
	Zoom              int       `json:"zoom"`
	UpdatedAt         time.Time `json:"updated_at"`
	CreatedAt         time.Time `json:"created_at"`
}
