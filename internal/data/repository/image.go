package repository

import (
	"github.com/google/uuid"

	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/core/model/dto"
	"github.com/Masedko/go-backend/internal/data/database"
)

type ImageRepo struct {
	db *database.DB
}

func NewImageRepository(db *database.DB) *ImageRepo {
	return &ImageRepo{
		db: db,
	}
}

func (r *ImageRepo) CreateImage(image model.Image) error {
	_, err := r.db.Exec(`
		INSERT INTO images (
			id,
			destroyed_object_id,
			file_name,
			path,
			lat,
			lng,
		    x,
		    y,
		    zoom,
			updated_at,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		image.ID, image.DestroyedObjectID, image.FileName, image.Path, image.Lat, image.Lng, image.X, image.Y, image.Zoom, image.UpdatedAt, image.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImageRepo) GetImageByLatLngXYZoom(lat float64, lng float64, x, y, zoom int) (dto.GetImage, error) {
	var image dto.GetImage
	err := r.db.Get(&image, `
		SELECT 
			id,
			destroyed_object_id,
			file_name,
			path,
			lat,
			lng,
			x,
			y,
			zoom,
			-- Haversine formula
			( acos( cos( radians($1) ) * cos( radians( lat ) ) * cos( radians( lng ) - radians($2) ) + sin( radians(:$1) ) * sin( radians( lat ) ) ) ) AS distance
		FROM images
		WHERE AND x = $3 AND y = $4 AND zoom = $5
		ORDER BY distance
		LIMIT 1`, lat, lng, x, y, zoom,
	)
	if err != nil {
		return dto.GetImage{}, err
	}

	return image, nil
}

func (r *ImageRepo) DeleteImage(id uuid.UUID) error {
	_, err := r.db.Exec(`
		DELETE FROM images
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImageRepo) DeleteImagesByDestroyedObjectID(destroyedObjectID uuid.UUID) error {
	_, err := r.db.Exec(`
		DELETE FROM images
		WHERE destroyed_object_id = $1`,
		destroyedObjectID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImageRepo) GetPathByID(id uuid.UUID) (string, error) {
	var path string
	err := r.db.Get(&path, `
		SELECT path
		FROM images
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return "", err
	}

	return path, nil
}

func (r *ImageRepo) GetPathsByDestroyedObjectID(destroyedObjectID uuid.UUID) ([]string, error) {
	var path []string
	err := r.db.Get(&path, `
		SELECT path
		FROM images
		WHERE destroyedObjectID = $1`,
		destroyedObjectID,
	)
	if err != nil {
		return nil, err
	}

	return path, nil
}
