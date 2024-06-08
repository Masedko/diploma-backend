package dto

import "github.com/google/uuid"

type GetImage struct {
	ID                uuid.UUID `json:"id"`
	DestroyedObjectID uuid.UUID `json:"destroyed_object_id"`
	FileName          string    `json:"file_name"`
	Path              string    `json:"path"`
	Lat               float64   `json:"lat"`
	Lng               float64   `json:"lng"`
	X                 int       `json:"x"`
	Y                 int       `json:"y"`
	Zoom              int       `json:"zoom"`
}

type GetImageRequest struct {
	Lat  float64 `json:"lat" validate:"required"`
	Lng  float64 `json:"lng" validate:"required"`
	X    int     `json:"x" validate:"required"`
	Y    int     `json:"y" validate:"required"`
	Zoom int     `json:"zoom" validate:"required"`
}

type CreateImageRequest struct {
	DestroyedObjectID uuid.UUID `json:"destroyed_object_id" validate:"required"`
	FileName          string    `json:"file_name" validate:"required"`
	Path              string    `json:"path" validate:"required"`
	Lat               float64   `json:"lat" validate:"required"`
	Lng               float64   `json:"lng" validate:"required"`
	X                 int       `json:"x" validate:"required"`
	Y                 int       `json:"y" validate:"required"`
	Zoom              int       `json:"zoom" validate:"required"`
}

type DeleteImageRequest struct {
	ID                *uuid.UUID `json:"id" validate:"optional"`
	DestroyedObjectID *uuid.UUID `json:"destroyed_object_id" validate:"optional"`
}

type GetImageResponse struct {
	ID                uuid.UUID `json:"id"`
	DestroyedObjectID uuid.UUID `json:"destroyed_object_id"`
	FileName          string    `json:"file_name"`
	Path              string    `json:"path"`
	Lat               float64   `json:"lat"`
	Lng               float64   `json:"lng"`
}

type CreateImageResponse struct{}

type DeleteImageResponse struct{}
