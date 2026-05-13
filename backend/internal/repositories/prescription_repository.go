package repositories

import (
	"context"
	"sterling-hms/internal/models"

	"github.com/jmoiron/sqlx"
)

type PrescriptionRepository struct {
	db *sqlx.DB
}

func NewPrescriptionRepository(db *sqlx.DB) *PrescriptionRepository {
	return &PrescriptionRepository{db: db}
}

func (r *PrescriptionRepository) Create(ctx context.Context, appointmentID, doctorID, patientID, diagnosis, notes string) (*models.Prescription, error) {
	var prescription models.Prescription
	query := `INSERT INTO prescriptions (appointment_id, doctor_id, patient_id, diagnosis, notes)
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id, appointment_id, doctor_id, patient_id, diagnosis, notes, dispensed, created_at`
	err := r.db.GetContext(ctx, &prescription, query, appointmentID, doctorID, patientID, diagnosis, notes)
	return &prescription, err
}

func (r *PrescriptionRepository) AddItem(ctx context.Context, prescriptionID, medicineID, dosage, frequency, instructions string, durationDays, quantity int) (*models.PrescriptionItem, error) {
	var item models.PrescriptionItem
	query := `INSERT INTO prescription_items (prescription_id, medicine_id, dosage, frequency, duration_days, quantity, instructions)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING id, prescription_id, medicine_id, dosage, frequency, duration_days, quantity, instructions`
	err := r.db.GetContext(ctx, &item, query, prescriptionID, medicineID, dosage, frequency, durationDays, quantity, instructions)
	return &item, err
}

func (r *PrescriptionRepository) GetByID(ctx context.Context, id string) (*models.Prescription, error) {
	var prescription models.Prescription
	query := `SELECT p.id, p.appointment_id, p.doctor_id, p.patient_id, p.diagnosis, p.notes, p.dispensed, p.created_at
			  FROM prescriptions p WHERE p.id = $1`
	err := r.db.GetContext(ctx, &prescription, query, id)
	return &prescription, err
}

func (r *PrescriptionRepository) GetItems(ctx context.Context, prescriptionID string) ([]models.PrescriptionItem, error) {
	var items []models.PrescriptionItem
	query := `SELECT id, prescription_id, medicine_id, dosage, frequency, duration_days, quantity, instructions
			  FROM prescription_items WHERE prescription_id = $1`
	err := r.db.SelectContext(ctx, &items, query, prescriptionID)
	return items, err
}

func (r *PrescriptionRepository) GetByPatientID(ctx context.Context, patientID string) ([]models.PrescriptionWithDetails, error) {
	var prescriptions []models.PrescriptionWithDetails
	query := `SELECT p.id, p.appointment_id, p.doctor_id, p.patient_id, p.diagnosis, p.notes, p.dispensed, p.created_at,
		  u.full_name as doctor_name
		  FROM prescriptions p
		  JOIN doctors d ON d.id = p.doctor_id
		  JOIN users u ON u.id = d.user_id
		  WHERE p.patient_id = $1 ORDER BY p.created_at DESC`
	err := r.db.SelectContext(ctx, &prescriptions, query, patientID)
	if err != nil {
		return nil, err
	}
	for i, p := range prescriptions {
		items, err := r.GetItems(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		prescriptions[i].Items = items
	}
	return prescriptions, nil
}

func (r *PrescriptionRepository) Dispense(ctx context.Context, id string) (*models.Prescription, error) {
	var prescription models.Prescription
	query := `UPDATE prescriptions SET dispensed = TRUE WHERE id = $1
			  RETURNING id, appointment_id, doctor_id, patient_id, diagnosis, notes, dispensed, created_at`
	err := r.db.GetContext(ctx, &prescription, query, id)
	return &prescription, err
}
