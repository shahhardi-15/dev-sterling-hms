package models

import "time"

type Patient struct {
	ID                    string    `db:"id" json:"id"`
	UserID                string    `db:"user_id" json:"user_id"`
	DOB                   *string   `db:"dob" json:"dob"`
	BloodGroup            *string   `db:"blood_group" json:"blood_group"`
	Phone                 *string   `db:"phone" json:"phone"`
	Address               *string   `db:"address" json:"address"`
	EmergencyContactName  *string   `db:"emergency_contact_name" json:"emergency_contact_name"`
	EmergencyContactPhone *string   `db:"emergency_contact_phone" json:"emergency_contact_phone"`
	CreatedAt             time.Time `db:"created_at" json:"created_at"`
}

type PatientWithUser struct {
	Patient
	FullName string `db:"full_name" json:"full_name"`
	Email    string `db:"email" json:"email"`
	IsActive bool   `db:"is_active" json:"is_active"`
}
