package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"

	pkgerrors "github.com/Masedko/go-backend/internal/core/errors"
	"github.com/Masedko/go-backend/internal/core/model"
)

func (h *Handler) GetDestroyedObjects(c echo.Context) error {
	req := &getDestroyedObjectsRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	objects, err := h.destroyedObjectRepo.GetDestroyedObjects(req.Page, req.PerPage, req.Name, req.Type, req.Region, req.DestructionTime, req.RestorationTime)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to get destroyed objects", err))
	}
	return c.JSON(http.StatusOK, getDestroyedObjectsResponse{
		DestroyedObjects: objects,
	})
}

func (h *Handler) CreateDestroyedObject(c echo.Context) error {
	req := &createDestroyObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}

	destructionTime, err := time.Parse(time.DateTime, req.DestructionTime)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to parse destruction date", err))
	}
	var restorationTime pq.NullTime
	if req.RestorationTime != nil {
		restorationTime.Time, err = time.Parse(time.DateTime, *req.RestorationTime)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to parse restoration date", err))
		}
		restorationTime.Valid = true
	}

	err = h.destroyedObjectRepo.CreateDestroyedObject(model.DestroyedObject{
		Name:            req.Name,
		Description:     req.Description,
		Type:            req.Type,
		Region:          req.Region,
		Address:         req.Address,
		Lat:             req.Lat,
		Lng:             req.Lng,
		DestructionTime: destructionTime,
		RestorationTime: restorationTime,
		UpdatedAt:       time.Now(),
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to create destroyed object", err))
	}
	return c.JSON(http.StatusCreated, createDestroyedObjectResponse{})
}

func (h *Handler) UpdateDestroyedObject(c echo.Context) error {
	req := &updateDestroyedObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}

	destructionTime, err := time.Parse(time.DateTime, req.DestructionTime)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to parse destruction date", err))
	}
	var restorationTime pq.NullTime
	if req.RestorationTime != nil {
		restorationTime.Time, err = time.Parse(time.DateTime, *req.RestorationTime)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to parse restoration date", err))
		}
		restorationTime.Valid = true
	}
	err = h.destroyedObjectRepo.UpdateDestroyedObject(model.DestroyedObject{
		ID:              req.ID,
		Name:            req.Name,
		Description:     req.Description,
		Type:            req.Type,
		Region:          req.Region,
		Address:         req.Address,
		Lat:             req.Lat,
		Lng:             req.Lng,
		DestructionTime: destructionTime,
		RestorationTime: restorationTime,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to update destroyed object", err))
	}
	return c.JSON(http.StatusCreated, updateDestroyedObjectResponse{})
}

func (h *Handler) DeleteDestroyedObject(c echo.Context) error {
	req := &deleteDestroyedObjectRequest{}
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to bind request", err))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to validate request", err))
	}
	err := h.destroyedObjectRepo.DeleteDestroyedObject(req.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, pkgerrors.NewError("Failed to delete destroyed object", err))
	}
	return c.JSON(http.StatusNoContent, deleteDestroyedObjectResponse{})
}
