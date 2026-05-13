package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type MedicineRepository struct {
	db *sqlx.DB
}

func NewMedicineRepository(db *sqlx.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}

func (r *MedicineRepository) GetAll(ctx context.Context) ([]models.Medicine, error) {
	var medicines []models.Medicine
	query := `SELECT id, name, generic_name, category, unit, price, reorder_level, is_active, created_at
			  FROM medicines WHERE is_active = true ORDER BY name`
	err := r.db.SelectContext(ctx, &medicines, query)
	return medicines, err
}

func (r *MedicineRepository) Create(ctx context.Context, name, genericName, category, unit string, price float64, reorderLevel int) (*models.Medicine, error) {
	var medicine models.Medicine
	query := `INSERT INTO medicines (name, generic_name, category, unit, price, reorder_level)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, name, generic_name, category, unit, price, reorder_level, is_active, created_at`
	err := r.db.GetContext(ctx, &medicine, query, name, genericName, category, unit, price, reorderLevel)
	return &medicine, err
}
