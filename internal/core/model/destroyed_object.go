package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type DestroyedObject struct {
	ID              uuid.UUID   `json:"id"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Type            string      `json:"type"`
	Region          string      `json:"region"`
	Address         string      `json:"address"`
	Lat             float64     `json:"lat"`
	Lng             float64     `json:"lng"`
	DestructionTime time.Time   `json:"destruction_time"`
	RestorationTime pq.NullTime `json:"restoration_time"`
	UpdatedAt       time.Time   `json:"updated_at"`
	CreatedAt       time.Time   `json:"created_at"`
}
