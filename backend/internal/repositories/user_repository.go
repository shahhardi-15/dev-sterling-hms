package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, full_name, email, role, is_active, created_at, updated_at FROM users ORDER BY created_at DESC`
	err := r.db.SelectContext(ctx, &users, query)
	return users, err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, full_name, email, role, is_active, created_at, updated_at FROM users WHERE id = $1`
	err := r.db.GetContext(ctx, &user, query, id)
	return &user, err
}

func (r *UserRepository) CreateUser(ctx context.Context, fullName, email, passwordHash, role string) (*models.User, error) {
	var user models.User
	query := `INSERT INTO users (full_name, email, password_hash, role) 
			  VALUES ($1, $2, $3, $4) 
			  RETURNING id, full_name, email, role, is_active, created_at, updated_at`
	err := r.db.GetContext(ctx, &user, query, fullName, email, passwordHash, role)
	return &user, err
}

func (r *UserRepository) UpdateUser(ctx context.Context, id, fullName string, isActive bool) (*models.User, error) {
	var user models.User
	query := `UPDATE users SET full_name = $1, is_active = $2, updated_at = NOW() 
			  WHERE id = $3 
			  RETURNING id, full_name, email, role, is_active, created_at, updated_at`
	err := r.db.GetContext(ctx, &user, query, fullName, isActive, id)
	return &user, err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id, passwordHash string) error {
        query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`
        _, err := r.db.ExecContext(ctx, query, passwordHash, id)
        return err
}

func (r *UserRepository) GetAllPatients(ctx context.Context) ([]models.PatientWithUser, error) {
	var patients []models.PatientWithUser
	query := `SELECT p.id, p.user_id, p.dob, p.blood_group, p.phone, p.address,
			  p.emergency_contact_name, p.emergency_contact_phone, p.created_at,
			  u.full_name, u.email, u.is_active
			  FROM patients p JOIN users u ON p.user_id = u.id
			  ORDER BY p.created_at DESC`
	err := r.db.SelectContext(ctx, &patients, query)
	return patients, err
}

func (r *UserRepository) GetAllDoctors(ctx context.Context) ([]models.DoctorWithUser, error) {
	var doctors []models.DoctorWithUser
	query := `SELECT d.id, d.user_id, d.department_id, d.specialization, d.license_no,
			  d.availability, d.bio, d.created_at,
			  u.full_name, u.email, u.is_active
			  FROM doctors d JOIN users u ON d.user_id = u.id
			  ORDER BY d.created_at DESC`
	err := r.db.SelectContext(ctx, &doctors, query)
	return doctors, err
}

func (r *UserRepository) GetDoctorByUserID(ctx context.Context, userID string) (*models.DoctorWithUser, error) {
	var doctor models.DoctorWithUser
	query := `SELECT d.id, d.user_id, d.department_id, d.specialization, d.license_no,
			  d.availability, d.bio, d.created_at,
			  u.full_name, u.email, u.is_active
			  FROM doctors d JOIN users u ON d.user_id = u.id
			  WHERE d.user_id = $1`
	err := r.db.GetContext(ctx, &doctor, query, userID)
	return &doctor, err
}
