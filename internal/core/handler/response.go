package handler

import "github.com/Masedko/go-backend/internal/core/model"

type getDestroyedObjectsResponse struct {
	DestroyedObjects []model.DestroyedObject `json:"destroyed_objects"`
}

type createDestroyedObjectResponse struct{}

type updateDestroyedObjectResponse struct{}

type deleteDestroyedObjectResponse struct{}

type getImagesResponse struct {
	Images []model.Image `json:"images"`
}

type createImageResponse struct{}

// updateImage is not implemented because of complications related with atomicity between storage and db

type deleteImageResponse struct{}
