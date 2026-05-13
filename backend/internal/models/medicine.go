package models

import "time"

type Medicine struct {
	ID           string    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	GenericName  *string   `db:"generic_name" json:"generic_name"`
	Category     *string   `db:"category" json:"category"`
	Unit         string    `db:"unit" json:"unit"`
	Price        float64   `db:"price" json:"price"`
	ReorderLevel int       `db:"reorder_level" json:"reorder_level"`
	IsActive     bool      `db:"is_active" json:"is_active"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
