package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/core/storage"
)

func (h *Handler) GetImages(c echo.Context) error {
	req := &getImagesRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	images, err := h.imagesRepo.GetImages(req.Page, req.PerPage)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to get images", err))
	}
	return c.JSON(http.StatusCreated, getImagesResponse{
		Images: images,
	})
}

func (h *Handler) CreateImage(c echo.Context) error {
	req := &createImageRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to get file", err))
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	path, err := h.client.UploadToBucket(c.Request().Context(), storage.ImageBucket,
		fmt.Sprintf("%f_%f/%d_%d_%d", req.Lat, req.Lng, req.X, req.Y, req.Zoom), src)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to upload file", err))
	}

	err = h.imagesRepo.CreateImage(model.Image{
		ID:                req.ID,
		DestroyedObjectID: req.DestroyedObjectID,
		FileName:          req.FileName,
		Path:              path,
		Lat:               req.Lat,
		Lng:               req.Lng,
		X:                 req.X,
		Y:                 req.Y,
		Zoom:              req.Zoom,
		UpdatedAt:         time.Now(),
		CreatedAt:         time.Now(),
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to create image", err))
	}
	return c.JSON(http.StatusCreated, createImageResponse{})
}

func (h *Handler) DeleteImage(c echo.Context) error {
	req := &deleteImageRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	if req.ID != nil {
		err := h.imagesRepo.DeleteImage(*req.ID)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to delete image", err))
		}
	} else if req.DestroyedObjectID != nil {
		err := h.imagesRepo.DeleteImagesByDestroyedObjectID(*req.DestroyedObjectID)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to delete images", err))
		}
	} else {
		return c.JSON(http.StatusBadRequest, pkgerrors.NewError("Failed to delete image", nil))
	}
	return c.JSON(http.StatusCreated, deleteImageResponse{})
}
