package models

import "time"

type Appointment struct {
	ID          string    `db:"id" json:"id"`
	PatientID   string    `db:"patient_id" json:"patient_id"`
	DoctorID    string    `db:"doctor_id" json:"doctor_id"`
	ScheduledAt time.Time `db:"scheduled_at" json:"scheduled_at"`
	Status      string    `db:"status" json:"status"`
	Type        string    `db:"type" json:"type"`
	Reason      *string   `db:"reason" json:"reason"`
	Notes       *string   `db:"notes" json:"notes"`
	CreatedBy   *string   `db:"created_by" json:"created_by"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type AppointmentWithDetails struct {
	Appointment
	PatientName    string  `db:"patient_name" json:"patient_name"`
	PatientEmail   string  `db:"patient_email" json:"patient_email"`
	DoctorName     string  `db:"doctor_name" json:"doctor_name"`
	Specialization *string `db:"specialization" json:"specialization"`
}
