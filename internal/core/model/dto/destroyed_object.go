package dto

import (
	"github.com/google/uuid"

	"github.com/Masedko/go-backend/internal/core/model"
)

type GetDestroyedObjectsRequest struct {
	Page                 int     `json:"page" validate:"required"`
	PerPage              int     `json:"per_page" validate:"required"`
	Name                 *string `json:"name" validate:"optional"`
	Type                 *string `json:"type" validate:"optional"`
	Region               *string `json:"region" validate:"optional"`
	StartDestructionTime *string `json:"start_destruction_time" validate:"optional"`
	EndDestructionTime   *string `json:"end_destruction_time" validate:"optional"`
	StartRestorationTime *string `json:"start_restoration_time" validate:"optional"`
	EndRestorationTime   *string `json:"end_restoration_time" validate:"optional"`
}

type CreateDestroyObjectRequest struct {
	Name            string  `json:"name" validate:"required"`
	Description     string  `json:"description" validate:"required"`
	Type            string  `json:"type" validate:"required"`
	Region          string  `json:"region" validate:"required"`
	Address         string  `json:"address" validate:"required"`
	Lat             float64 `json:"lat" validate:"required"`
	Lng             float64 `json:"lng" validate:"required"`
	DestructionTime string  `json:"destruction_time" validate:"required"`
	RestorationTime *string `json:"restoration_time" validate:"optional"`
}

type UpdateDestroyedObjectRequest struct {
	ID              uuid.UUID `json:"id" validate:"required"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	Type            string    `json:"type" validate:"required"`
	Region          string    `json:"region" validate:"required"`
	Address         string    `json:"address" validate:"required"`
	Lat             float64   `json:"lat" validate:"required"`
	Lng             float64   `json:"lng" validate:"required"`
	DestructionTime string    `json:"destruction_time" validate:"required"`
	RestorationTime *string   `json:"restoration_time" validate:"optional"`
}

type DeleteDestroyedObjectRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type GetDestroyedObjectsResponse struct {
	DestroyedObjects []model.DestroyedObject `json:"destroyed_objects"`
}

type CreateDestroyedObjectResponse struct{}

type UpdateDestroyedObjectResponse struct{}

type DeleteDestroyedObjectResponse struct{}
