package models

import "time"

type Prescription struct {
	ID            string    `db:"id" json:"id"`
	AppointmentID string    `db:"appointment_id" json:"appointment_id"`
	DoctorID      string    `db:"doctor_id" json:"doctor_id"`
	PatientID     string    `db:"patient_id" json:"patient_id"`
	Diagnosis     string    `db:"diagnosis" json:"diagnosis"`
	Notes         *string   `db:"notes" json:"notes"`
	Dispensed     bool      `db:"dispensed" json:"dispensed"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type PrescriptionItem struct {
	ID             string  `db:"id" json:"id"`
	PrescriptionID string  `db:"prescription_id" json:"prescription_id"`
	MedicineID     string  `db:"medicine_id" json:"medicine_id"`
	Dosage         *string `db:"dosage" json:"dosage"`
	Frequency      *string `db:"frequency" json:"frequency"`
	DurationDays   *int    `db:"duration_days" json:"duration_days"`
	Quantity       int     `db:"quantity" json:"quantity"`
	Instructions   *string `db:"instructions" json:"instructions"`
}

type PrescriptionWithDetails struct {
	Prescription
	PatientName string             `db:"patient_name" json:"patient_name"`
	DoctorName  string             `db:"doctor_name" json:"doctor_name"`
	Items       []PrescriptionItem `json:"items"`
}
