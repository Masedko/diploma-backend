package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/core/model/dto"
)

type DestroyedObjectRepo interface {
	GetDestroyedObjects(
		page int, perPage int,
		name *string,
		objType *string,
		region *string,
		startDestructionTime *string, endDestructionTime *string,
		startRestorationTime *string, endRestorationTime *string,
	) ([]model.DestroyedObject, error)
	CreateDestroyedObject(object model.DestroyedObject) error
	UpdateDestroyedObject(object model.DestroyedObject) error
	DeleteDestroyedObject(id uuid.UUID) error
}

type DestroyedObjectService struct {
	destroyedObjectRepo DestroyedObjectRepo
}

func NewDestroyedObjectService(destroyedObjectRepo DestroyedObjectRepo) *DestroyedObjectService {
	return &DestroyedObjectService{
		destroyedObjectRepo: destroyedObjectRepo,
	}
}

func (s *DestroyedObjectService) GetDestroyedObjects(req *dto.GetDestroyedObjectsRequest) ([]model.DestroyedObject, error) {
	return s.destroyedObjectRepo.GetDestroyedObjects(
		req.Page, req.PerPage,
		req.Name, req.Type, req.Region,
		req.StartDestructionTime, req.EndDestructionTime,
		req.StartRestorationTime, req.EndRestorationTime,
	)
}

func (s *DestroyedObjectService) CreateDestroyedObject(req *dto.CreateDestroyObjectRequest) error {
	destructionTime, err := time.Parse(time.DateTime, req.DestructionTime)
	if err != nil {
		return err
	}
	var restorationTime pq.NullTime
	if req.RestorationTime != nil {
		restorationTime.Time, err = time.Parse(time.DateTime, *req.RestorationTime)
		if err != nil {
			return err
		}
		restorationTime.Valid = true
	}
	return s.destroyedObjectRepo.CreateDestroyedObject(model.DestroyedObject{
		ID:              uuid.New(),
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
}

func (s *DestroyedObjectService) UpdateDestroyedObject(req *dto.UpdateDestroyedObjectRequest) error {
	destructionTime, err := time.Parse(time.DateTime, req.DestructionTime)
	if err != nil {
		return err
	}
	var restorationTime pq.NullTime
	if req.RestorationTime != nil {
		restorationTime.Time, err = time.Parse(time.DateTime, *req.RestorationTime)
		if err != nil {
			return err
		}
		restorationTime.Valid = true
	}
	return s.destroyedObjectRepo.UpdateDestroyedObject(model.DestroyedObject{
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
		UpdatedAt:       time.Now(),
	})
}

func (s *DestroyedObjectService) DeleteDestroyedObject(req *dto.DeleteDestroyedObjectRequest) error {
	return s.destroyedObjectRepo.DeleteDestroyedObject(req.ID)
}
