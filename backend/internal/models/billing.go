package models

import "time"

type Billing struct {
	ID              string     `db:"id" json:"id"`
	AppointmentID   string     `db:"appointment_id" json:"appointment_id"`
	PatientID       string     `db:"patient_id" json:"patient_id"`
	ConsultationFee float64    `db:"consultation_fee" json:"consultation_fee"`
	MedicineCost    float64    `db:"medicine_cost" json:"medicine_cost"`
	OtherCharges    float64    `db:"other_charges" json:"other_charges"`
	Discount        float64    `db:"discount" json:"discount"`
	TotalAmount     float64    `db:"total_amount" json:"total_amount"`
	PaymentStatus   string     `db:"payment_status" json:"payment_status"`
	PaymentMethod   *string    `db:"payment_method" json:"payment_method"`
	PaidAt          *time.Time `db:"paid_at" json:"paid_at"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
}

type BillingWithDetails struct {
	Billing
	PatientName  string `db:"patient_name" json:"patient_name"`
	PatientEmail string `db:"patient_email" json:"patient_email"`
}
