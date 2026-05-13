package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, full_name, email, password_hash, role, is_active, created_at, updated_at 
			  FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, &user, query, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) FindUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, full_name, email, password_hash, role, is_active, created_at, updated_at 
			  FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *AuthRepository) UpdatePassword(ctx context.Context, userID, hashedPassword string) error {
	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, hashedPassword, userID)
	return err
}
