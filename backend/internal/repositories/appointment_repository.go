package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository struct {
	db *sqlx.DB
}

func NewAppointmentRepository(db *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) GetAll(ctx context.Context) ([]models.AppointmentWithDetails, error) {
	var appointments []models.AppointmentWithDetails
	query := `SELECT a.id, a.patient_id, a.doctor_id, a.scheduled_at, a.status, a.type,
			  a.reason, a.notes, a.created_by, a.created_at, a.updated_at,
			  u1.full_name as patient_name, u1.email as patient_email,
			  u2.full_name as doctor_name, d.specialization
			  FROM appointments a
			  JOIN patients p ON a.patient_id = p.id
			  JOIN users u1 ON p.user_id = u1.id
			  JOIN doctors d ON a.doctor_id = d.id
			  JOIN users u2 ON d.user_id = u2.id
			  ORDER BY a.scheduled_at DESC`
	err := r.db.SelectContext(ctx, &appointments, query)
	return appointments, err
}

func (r *AppointmentRepository) GetByID(ctx context.Context, id string) (*models.AppointmentWithDetails, error) {
	var appointment models.AppointmentWithDetails
	query := `SELECT a.id, a.patient_id, a.doctor_id, a.scheduled_at, a.status, a.type,
			  a.reason, a.notes, a.created_by, a.created_at, a.updated_at,
			  u1.full_name as patient_name, u1.email as patient_email,
			  u2.full_name as doctor_name, d.specialization
			  FROM appointments a
			  JOIN patients p ON a.patient_id = p.id
			  JOIN users u1 ON p.user_id = u1.id
			  JOIN doctors d ON a.doctor_id = d.id
			  JOIN users u2 ON d.user_id = u2.id
			  WHERE a.id = $1`
	err := r.db.GetContext(ctx, &appointment, query, id)
	return &appointment, err
}

func (r *AppointmentRepository) GetByPatientID(ctx context.Context, patientID string) ([]models.AppointmentWithDetails, error) {
	var appointments []models.AppointmentWithDetails
	query := `SELECT a.id, a.patient_id, a.doctor_id, a.scheduled_at, a.status, a.type,
			  a.reason, a.notes, a.created_by, a.created_at, a.updated_at,
			  u1.full_name as patient_name, u1.email as patient_email,
			  u2.full_name as doctor_name, d.specialization
			  FROM appointments a
			  JOIN patients p ON a.patient_id = p.id
			  JOIN users u1 ON p.user_id = u1.id
			  JOIN doctors d ON a.doctor_id = d.id
			  JOIN users u2 ON d.user_id = u2.id
			  WHERE a.patient_id = $1
			  ORDER BY a.scheduled_at DESC`
	err := r.db.SelectContext(ctx, &appointments, query, patientID)
	return appointments, err
}

func (r *AppointmentRepository) GetByDoctorID(ctx context.Context, doctorID string) ([]models.AppointmentWithDetails, error) {
	var appointments []models.AppointmentWithDetails
	query := `SELECT a.id, a.patient_id, a.doctor_id, a.scheduled_at, a.status, a.type,
			  a.reason, a.notes, a.created_by, a.created_at, a.updated_at,
			  u1.full_name as patient_name, u1.email as patient_email,
			  u2.full_name as doctor_name, d.specialization
			  FROM appointments a
			  JOIN patients p ON a.patient_id = p.id
			  JOIN users u1 ON p.user_id = u1.id
			  JOIN doctors d ON a.doctor_id = d.id
			  JOIN users u2 ON d.user_id = u2.id
			  WHERE a.doctor_id = $1
			  ORDER BY a.scheduled_at DESC`
	err := r.db.SelectContext(ctx, &appointments, query, doctorID)
	return appointments, err
}

func (r *AppointmentRepository) Create(ctx context.Context, patientID, doctorID, scheduledAt, apptType, reason, createdBy string) (*models.Appointment, error) {
	var appointment models.Appointment
	query := `INSERT INTO appointments (patient_id, doctor_id, scheduled_at, type, reason, created_by)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, patient_id, doctor_id, scheduled_at, status, type, reason, notes, created_by, created_at, updated_at`
	err := r.db.GetContext(ctx, &appointment, query, patientID, doctorID, scheduledAt, apptType, reason, createdBy)
	return &appointment, err
}

func (r *AppointmentRepository) UpdateStatus(ctx context.Context, id, status string) (*models.Appointment, error) {
	var appointment models.Appointment
	query := `UPDATE appointments SET status = $1, updated_at = NOW()
			  WHERE id = $2
			  RETURNING id, patient_id, doctor_id, scheduled_at, status, type, reason, notes, created_by, created_at, updated_at`
	err := r.db.GetContext(ctx, &appointment, query, status, id)
	return &appointment, err
}
