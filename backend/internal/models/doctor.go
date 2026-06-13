package models

import (
	"encoding/json"
	"time"
)

type Doctor struct {
	ID             string          `db:"id" json:"id"`
	UserID         string          `db:"user_id" json:"user_id"`
	DepartmentID   *string         `db:"department_id" json:"department_id"`
	Specialization *string         `db:"specialization" json:"specialization"`
	LicenseNo      *string         `db:"license_no" json:"license_no"`
	Availability   json.RawMessage `db:"availability" json:"availability"`
	Bio            *string         `db:"bio" json:"bio"`
	CreatedAt      time.Time       `db:"created_at" json:"created_at"`
}

type DoctorWithUser struct {
	Doctor
	FullName string `db:"full_name" json:"full_name"`
	Email    string `db:"email" json:"email"`
	IsActive bool   `db:"is_active" json:"is_active"`
}