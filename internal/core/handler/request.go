package handler

import "github.com/google/uuid"

type getDestroyedObjectsRequest struct {
	Page            int     `json:"page" validate:"required"`
	PerPage         int     `json:"per_page" validate:"required"`
	Name            *string `json:"name" validate:"required"`
	Type            *string `json:"type" validate:"required"`
	Region          *string `json:"region" validate:"required"`
	DestructionTime *string `json:"destruction_time" validate:"required"`
	RestorationTime *string `json:"restoration_time" validate:"optional"`
}

type createDestroyObjectRequest struct {
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

type updateDestroyedObjectRequest struct {
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

type deleteDestroyedObjectRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type getImagesRequest struct {
	Page              int        `json:"page" validate:"required"`
	PerPage           int        `json:"per_page" validate:"required"`
	DestroyedObjectID *uuid.UUID `json:"destroyed_object_id" validate:"optional"`
}

type createImageRequest struct {
	ID                uuid.UUID `json:"id" validate:"required"`
	DestroyedObjectID uuid.UUID `json:"destroyed_object_id" validate:"required"`
	FileName          string    `json:"file_name" validate:"required"`
	Path              string    `json:"path" validate:"required"`
	Lat               float64   `json:"lat" validate:"required"`
	Lng               float64   `json:"lng" validate:"required"`
	X                 int       `json:"x" validate:"required"`
	Y                 int       `json:"y" validate:"required"`
	Zoom              int       `json:"zoom" validate:"required"`
}

type deleteImageRequest struct {
	ID                *uuid.UUID `json:"id" validate:"optional"`
	DestroyedObjectID *uuid.UUID `json:"destroyed_object_id" validate:"optional"`
}
