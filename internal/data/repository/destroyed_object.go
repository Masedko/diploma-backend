package repository

import (
	"github.com/google/uuid"

	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/data/database"
)

type DestroyedObjectRepo struct {
	db *database.DB
}

func NewDestroyedObjectRepository(db *database.DB) *DestroyedObjectRepo {
	return &DestroyedObjectRepo{
		db: db,
	}
}

func (r *DestroyedObjectRepo) GetDestroyedObjects(
	page int,
	perPage int,
	name *string,
	objType *string,
	region *string,
	startDestructionTime *string,
	endDestructionTime *string,
	startRestorationTime *string,
	endRestorationTime *string,
) ([]model.DestroyedObject, error) {
	var objects []model.DestroyedObject
	var conditions []string
	var args []interface{}
	var argID int

	argID++
	conditions = append(conditions, "LIMIT $"+string(rune(argID)))
	args = append(args, perPage)

	argID++
	conditions = append(conditions, "OFFSET $"+string(rune(argID)))
	args = append(args, (page-1)*perPage)

	if name != nil {
		argID++
		conditions = append(conditions, "name = $"+string(rune(argID)))
		args = append(args, *name)
	}

	if objType != nil {
		argID++
		conditions = append(conditions, "type = $"+string(rune(argID)))
		args = append(args, *objType)
	}

	if region != nil {
		argID++
		conditions = append(conditions, "region = $"+string(rune(argID)))
		args = append(args, *region)
	}

	if startDestructionTime != nil {
		argID++
		conditions = append(conditions, "destruction_time >= $"+string(rune(argID)))
		args = append(args, *startDestructionTime)
	}

	if endDestructionTime != nil {
		argID++
		conditions = append(conditions, "destruction_time <= $"+string(rune(argID)))
		args = append(args, *endDestructionTime)
	}

	if startRestorationTime != nil {
		argID++
		conditions = append(conditions, "restoration_time >= $"+string(rune(argID)))
		args = append(args, *startRestorationTime)
	}

	if endRestorationTime != nil {
		argID++
		conditions = append(conditions, "restoration_time <= $"+string(rune(argID)))
		args = append(args, *endRestorationTime)
	}

	baseQuery := `
		SELECT 
			id,
			name,
			description,
			type,
			region,
			address,
			lat,
			lng,
			destruction_time,
			restoration_time,
			updated_at,
			created_at,
		FROM destroyed_objects`
	err := r.db.Select(&objects, baseQuery)
	if err != nil {
		return nil, err
	}

	return objects, nil
}

func (r *DestroyedObjectRepo) CreateDestroyedObject(object model.DestroyedObject) error {
	_, err := r.db.Exec(`
		INSERT INTO destroyed_objects (
			id,
			name,
			description,
			type,
			region,
			address,
			lat,
			lng,
			destruction_time,
			restoration_time,
			updated_at,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		object.ID,
		object.Name,
		object.Description,
		object.Type,
		object.Region,
		object.Address,
		object.Lat,
		object.Lng,
		object.DestructionTime,
		object.RestorationTime,
		object.UpdatedAt,
		object.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *DestroyedObjectRepo) UpdateDestroyedObject(object model.DestroyedObject) error {
	_, err := r.db.Exec(`
		UPDATE destroyed_objects SET
			name = $1,
			description = $2,
			type = $3,
			region = $4,
			address = $5,
			lat = $6,
			lng = $7,
			destruction_time = $8,
			restoration_time = $9,
			updated_at = $10
		WHERE id = $11`,
		object.Name,
		object.Description,
		object.Type,
		object.Region,
		object.Address,
		object.Lat,
		object.Lng,
		object.DestructionTime,
		object.RestorationTime,
		object.UpdatedAt,
		object.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *DestroyedObjectRepo) DeleteDestroyedObject(id uuid.UUID) error {
	_, err := r.db.Exec(`
		DELETE FROM destroyed_objects
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
