package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/core/model/dto"
	"github.com/Masedko/go-backend/internal/core/service"
)

type DestroyedObjectService interface {
	GetDestroyedObjects(req *dto.GetDestroyedObjectsRequest) ([]model.DestroyedObject, error)
	CreateDestroyedObject(req *dto.CreateDestroyObjectRequest) error
	UpdateDestroyedObject(req *dto.UpdateDestroyedObjectRequest) error
	DeleteDestroyedObject(req *dto.DeleteDestroyedObjectRequest) error
}

type DestroyedObjectHandler struct {
	destroyedObjectService *service.DestroyedObjectService
}

func NewDestroyedObjectHandler(destroyedObjectService *service.DestroyedObjectService) *DestroyedObjectHandler {
	return &DestroyedObjectHandler{
		destroyedObjectService: destroyedObjectService,
	}
}

func (h *DestroyedObjectHandler) GetDestroyedObjects(c echo.Context) error {
	req := &dto.GetDestroyedObjectsRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	objects, err := h.destroyedObjectService.GetDestroyedObjects(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to get destroyed objects", err))
	}
	return c.JSON(http.StatusOK, dto.GetDestroyedObjectsResponse{
		DestroyedObjects: objects,
	})
}

func (h *DestroyedObjectHandler) CreateDestroyedObject(c echo.Context) error {
	req := &dto.CreateDestroyObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	err := h.destroyedObjectService.CreateDestroyedObject(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to create destroyed object", err))
	}
	return c.JSON(http.StatusCreated, dto.CreateDestroyedObjectResponse{})
}

func (h *DestroyedObjectHandler) UpdateDestroyedObject(c echo.Context) error {
	req := &dto.UpdateDestroyedObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	err := h.destroyedObjectService.UpdateDestroyedObject(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to update destroyed object", err))
	}
	return c.JSON(http.StatusNoContent, dto.UpdateDestroyedObjectResponse{})
}

func (h *DestroyedObjectHandler) DeleteDestroyedObject(c echo.Context) error {
	req := &dto.DeleteDestroyedObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	err := h.destroyedObjectService.DeleteDestroyedObject(req)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to delete destroyed object", err))
	}
	return c.JSON(http.StatusNoContent, dto.DeleteDestroyedObjectResponse{})
}
