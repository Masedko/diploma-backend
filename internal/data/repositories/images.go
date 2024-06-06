package repositories

import (
	"github.com/google/uuid"

	"github.com/Masedko/go-backend/internal/core/model"
	"github.com/Masedko/go-backend/internal/data/database"
)

type ImagesRepo struct {
	db *database.DB
}

func NewImagesRepository(db *database.DB) *ImagesRepo {
	return &ImagesRepo{
		db: db,
	}
}

func (r *ImagesRepo) GetImages(page int, perPage int) ([]model.Image, error) {
	var images []model.Image
	err := r.db.Select(&images, `
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
			updated_at,
			created_at
		FROM images
		LIMIT $1 OFFSET $2`,
		perPage, (page-1)*perPage,
	)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ImagesRepo) CreateImage(image model.Image) error {
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

func (r *ImagesRepo) GetImagesByDestroyedObjectID(destroyedObjectID uuid.UUID) ([]model.Image, error) {
	var images []model.Image
	err := r.db.Select(&images, `
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
			updated_at,
			created_at
		FROM images
		WHERE destroyed_object_id = $1`,
		destroyedObjectID,
	)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ImagesRepo) GetImageByDestroyedObjectIDLatLng(destroyedObjectID uuid.UUID, lat float64, lng float64) (*model.Image, error) {
	var image model.Image
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
			updated_at,
			created_at,
			-- Haversine formula
			( acos( cos( radians($2) ) * cos( radians( lat ) ) * cos( radians( lng ) - radians($3) ) + sin( radians(:$2) ) * sin( radians( lat ) ) ) ) AS distance
		FROM images
		WHERE destroyed_object_id = $1
		ORDER BY distance
		LIMIT 1`,
		destroyedObjectID, lat, lng,
	)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *ImagesRepo) UpdateImage(image model.Image) error {
	_, err := r.db.Exec(`
		UPDATE images
		SET
			file_name = $1,
			path = $2,
			lat = $3,
			lng = $4,
			x = $5,
			y = $6,
			zoom = $7,
			updated_at = $8
		WHERE id = $9`,
		image.FileName, image.Path, image.Lat, image.Lng, image.X, image.Y, image.Zoom, image.UpdatedAt, image.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *ImagesRepo) DeleteImage(id uuid.UUID) error {
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

func (r *ImagesRepo) DeleteImagesByDestroyedObjectID(destroyedObjectID uuid.UUID) error {
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
