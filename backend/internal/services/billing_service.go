package services

import (
	"context"
	"errors"
	"sterling-hms/internal/models"
	"sterling-hms/internal/repositories"
)

type BillingService struct {
	repo *repositories.BillingRepository
}

func NewBillingService(repo *repositories.BillingRepository) *BillingService {
	return &BillingService{repo: repo}
}

type CreateBillingRequest struct {
	AppointmentID   string  `json:"appointment_id" binding:"required"`
	PatientID       string  `json:"patient_id" binding:"required"`
	ConsultationFee float64 `json:"consultation_fee"`
	MedicineCost    float64 `json:"medicine_cost"`
	OtherCharges    float64 `json:"other_charges"`
	Discount        float64 `json:"discount"`
}

type MarkPaidRequest struct {
	PaymentMethod string `json:"payment_method" binding:"required"`
}

func (s *BillingService) Create(ctx context.Context, req CreateBillingRequest) (*models.Billing, error) {
	billing, err := s.repo.Create(ctx, req.AppointmentID, req.PatientID, req.ConsultationFee, req.MedicineCost, req.OtherCharges, req.Discount)
	if err != nil {
		return nil, errors.New("failed to create billing record")
	}
	return billing, nil
}

func (s *BillingService) GetByID(ctx context.Context, id string) (*models.BillingWithDetails, error) {
	billing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("billing record not found")
	}
	return billing, nil
}

func (s *BillingService) GetAll(ctx context.Context) ([]models.BillingWithDetails, error) {
	return s.repo.GetAll(ctx)
}

func (s *BillingService) MarkAsPaid(ctx context.Context, id string, req MarkPaidRequest) (*models.Billing, error) {
	billing, err := s.repo.MarkAsPaid(ctx, id, req.PaymentMethod)
	if err != nil {
		return nil, errors.New("failed to mark billing as paid")
	}
	return billing, nil
}
