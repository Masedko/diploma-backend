package handler

import (
	"github.com/Masedko/go-backend/internal/core/storage"
	"github.com/Masedko/go-backend/internal/data/database"
	"github.com/Masedko/go-backend/internal/data/repositories"
)

type Handler struct {
	destroyedObjectRepo *repositories.DestroyedObjectRepo
	imagesRepo          *repositories.ImagesRepo
	client              *storage.Client // Can be abstracted
}

func New(db *database.DB, s *storage.Client) *Handler {
	return &Handler{
		destroyedObjectRepo: repositories.NewDestroyedObjectRepository(db),
		imagesRepo:          repositories.NewImagesRepository(db),
		client:              s,
	}
}
