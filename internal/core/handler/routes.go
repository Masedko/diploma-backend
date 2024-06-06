package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	destroyedObjects := v1.Group("/destroyed_objects")
	destroyedObjects.GET("", h.GetDestroyedObjects)
	destroyedObjects.POST("", h.CreateDestroyedObject)
	destroyedObjects.PUT("", h.UpdateDestroyedObject)
	destroyedObjects.DELETE("", h.DeleteDestroyedObject)

	images := destroyedObjects.Group("/images")
	images.GET("", h.GetImages)
	images.POST("", h.CreateImage)
	images.DELETE("", h.DeleteImage)
}
