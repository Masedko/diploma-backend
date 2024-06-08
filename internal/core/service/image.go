package service

import (
	"context"
	"io"
	"mime/multipart"

	"github.com/google/uuid"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/core/model/dto"
)

type ImageRepo interface {
	CreateImage(image model.Image) error
	GetImageByLatLngXYZoom(lat float64, lng float64, x int, y int, zoom int) (dto.GetImage, error)
	DeleteImage(id uuid.UUID) error
	DeleteImagesByDestroyedObjectID(destroyedObjectID uuid.UUID) error
	GetPathByID(id uuid.UUID) (string, error)
	GetPathsByDestroyedObjectID(destroyedObjectID uuid.UUID) ([]string, error)
}

type ImageBucket interface {
	Upload(ctx context.Context, fileName string, file io.Reader) (string, error)
	Delete(ctx context.Context, fileName string) error
}

type ImageService struct {
	imageRepo   ImageRepo
	imageBucket ImageBucket
}

func NewImageService(imageRepo ImageRepo, imageBucket ImageBucket) *ImageService {
	return &ImageService{
		imageRepo:   imageRepo,
		imageBucket: imageBucket,
	}
}

func (s *ImageService) CreateImage(req *dto.CreateImageRequest, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	return s.imageRepo.CreateImage(model.Image{
		ID:                uuid.New(),
		Lat:               req.Lat,
		Lng:               req.Lng,
		X:                 req.X,
		Y:                 req.Y,
		Zoom:              req.Zoom,
		DestroyedObjectID: req.DestroyedObjectID,
	})
}

func (s *ImageService) GetImageByLatLngXYZoom(req *dto.GetImageRequest) (dto.GetImage, error) {
	return s.imageRepo.GetImageByLatLngXYZoom(req.Lat, req.Lng, req.X, req.Y, req.Zoom)
}

func (s *ImageService) DeleteImage(req *dto.DeleteImageRequest) error {
	if req.ID != nil {
		path, err := s.imageRepo.GetPathByID(*req.ID)
		if err != nil {
			return err
		}

		err = s.imageBucket.Delete(context.Background(), path)
		if err != nil {
			return err
		}
		err = s.imageRepo.DeleteImage(*req.ID)
		if err != nil {
			return err
		}
	} else if req.DestroyedObjectID != nil {
		imagePaths, err := s.imageRepo.GetPathsByDestroyedObjectID(*req.DestroyedObjectID)
		if err != nil {
			return err
		}
		for _, imagePath := range imagePaths {
			err = s.imageBucket.Delete(context.Background(), imagePath)
			if err != nil {
				return err
			}
		}
		err = s.imageRepo.DeleteImagesByDestroyedObjectID(*req.DestroyedObjectID)
		if err != nil {
			return err
		}
	} else {
		return pkgerrors.NewError("Specify ID or DestroyedObjectID", nil)
	}
	return nil
}
