package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type BillingRepository struct {
	db *sqlx.DB
}

func NewBillingRepository(db *sqlx.DB) *BillingRepository {
	return &BillingRepository{db: db}
}

func (r *BillingRepository) Create(ctx context.Context, appointmentID, patientID string, consultationFee, medicineCost, otherCharges, discount float64) (*models.Billing, error) {
	var billing models.Billing
	query := `INSERT INTO billing (appointment_id, patient_id, consultation_fee, medicine_cost, other_charges, discount)
			  VALUES ($1, $2, $3, $4, $5, $6)
			  RETURNING id, appointment_id, patient_id, consultation_fee, medicine_cost, other_charges, discount, total_amount, payment_status, payment_method, paid_at, created_at`
	err := r.db.GetContext(ctx, &billing, query, appointmentID, patientID, consultationFee, medicineCost, otherCharges, discount)
	return &billing, err
}

func (r *BillingRepository) GetByID(ctx context.Context, id string) (*models.BillingWithDetails, error) {
	var billing models.BillingWithDetails
	query := `SELECT b.id, b.appointment_id, b.patient_id, b.consultation_fee, b.medicine_cost,
			  b.other_charges, b.discount, b.total_amount, b.payment_status, b.payment_method, b.paid_at, b.created_at,
			  u.full_name as patient_name, u.email as patient_email
			  FROM billing b
			  JOIN patients p ON b.patient_id = p.id
			  JOIN users u ON p.user_id = u.id
			  WHERE b.id = $1`
	err := r.db.GetContext(ctx, &billing, query, id)
	return &billing, err
}

func (r *BillingRepository) GetAll(ctx context.Context) ([]models.BillingWithDetails, error) {
	var billings []models.BillingWithDetails
	query := `SELECT b.id, b.appointment_id, b.patient_id, b.consultation_fee, b.medicine_cost,
			  b.other_charges, b.discount, b.total_amount, b.payment_status, b.payment_method, b.paid_at, b.created_at,
			  u.full_name as patient_name, u.email as patient_email
			  FROM billing b
			  JOIN patients p ON b.patient_id = p.id
			  JOIN users u ON p.user_id = u.id
			  ORDER BY b.created_at DESC`
	err := r.db.SelectContext(ctx, &billings, query)
	return billings, err
}

func (r *BillingRepository) MarkAsPaid(ctx context.Context, id, paymentMethod string) (*models.Billing, error) {
	var billing models.Billing
	query := `UPDATE billing SET payment_status = 'paid', payment_method = $1, paid_at = NOW()
			  WHERE id = $2
			  RETURNING id, appointment_id, patient_id, consultation_fee, medicine_cost, other_charges, discount, total_amount, payment_status, payment_method, paid_at, created_at`
	err := r.db.GetContext(ctx, &billing, query, paymentMethod, id)
	return &billing, err
}

func (r *BillingRepository) CreateSimple(ctx context.Context, appointmentID, patientID string, amount float64) error {
	query := `INSERT INTO billing (appointment_id, patient_id, amount, status)
			  VALUES ($1, $2, $3, 'pending')
			  ON CONFLICT DO NOTHING`
	_, err := r.db.ExecContext(ctx, query, appointmentID, patientID, amount)
	return err
}