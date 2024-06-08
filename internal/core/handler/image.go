package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/model/dto"
	"github.com/Masedko/go-backend/internal/core/service"
)

type ImageHandler struct {
	imageService *service.ImageService
}

func NewImageHandler(imageService *service.ImageService) *ImageHandler {
	return &ImageHandler{
		imageService: imageService,
	}
}

func (h *ImageHandler) GetImage(c echo.Context) error {
	req := &dto.GetImageRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	dtoImageGet, err := h.imageService.GetImageByLatLngXYZoom(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to get images", err))
	}

	return c.JSON(http.StatusOK, dto.GetImageResponse{
		ID:                dtoImageGet.ID,
		DestroyedObjectID: dtoImageGet.DestroyedObjectID,
		FileName:          dtoImageGet.FileName,
		Path:              dtoImageGet.Path,
		Lat:               dtoImageGet.Lat,
		Lng:               dtoImageGet.Lng,
	})
}

func (h *ImageHandler) CreateImage(c echo.Context) error {
	req := &dto.CreateImageRequest{}
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

	err = h.imageService.CreateImage(req, file)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to create image", err))
	}

	return c.JSON(http.StatusCreated, dto.CreateImageResponse{})
}

func (h *ImageHandler) DeleteImage(c echo.Context) error {
	req := &dto.DeleteImageRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	err := h.imageService.DeleteImage(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to delete image", err))
	}

	return c.JSON(http.StatusNoContent, dto.DeleteImageResponse{})
}
