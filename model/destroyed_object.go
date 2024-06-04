package model

import (
	"time"

	"github.com/google/uuid"
)

type DestroyedObject struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Type            string    `json:"type"`
	Region          string    `json:"region"`
	Address         string    `json:"address"`
	DestructionDate string    `json:"destruction_date"`
	RestorationDate *string   `json:"restoration_date"`
	Images          []Image   `json:"images"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
}

type Image struct {
	ID        uuid.UUID `json:"id"`
	ObjectID  uuid.UUID `json:"object_id"`
	FileName  string    `json:"file_name"`
	Path      string    `json:"path"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
