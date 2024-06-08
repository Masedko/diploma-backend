package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *DestroyedObjectHandler) Register(v1 *echo.Group) {
	destroyedObjects := v1.Group("/destroyed_objects")
	destroyedObjects.GET("", h.GetDestroyedObjects)
	destroyedObjects.POST("", h.CreateDestroyedObject)
	destroyedObjects.PUT("", h.UpdateDestroyedObject)
	destroyedObjects.DELETE("", h.DeleteDestroyedObject)
}

func (h *ImageHandler) Register(v1 *echo.Group) {
	images := v1.Group("/images")
	images.GET("", h.GetImage)
	images.POST("", h.CreateImage)
	images.DELETE("", h.DeleteImage)
}
