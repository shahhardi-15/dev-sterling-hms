package services

import (
	"context"
	"errors"
	"sterling-hms/internal/models"
	"sterling-hms/internal/repositories"
)

type PrescriptionService struct {
	repo        *repositories.PrescriptionRepository
	billingRepo *repositories.BillingRepository
}

func NewPrescriptionService(repo *repositories.PrescriptionRepository, billingRepo *repositories.BillingRepository) *PrescriptionService {
	return &PrescriptionService{repo: repo, billingRepo: billingRepo}
}

type CreatePrescriptionRequest struct {
	AppointmentID string                    `json:"appointment_id" binding:"required"`
	DoctorID      string                    `json:"doctor_id" binding:"required"`
	PatientID     string                    `json:"patient_id" binding:"required"`
	Diagnosis     string                    `json:"diagnosis" binding:"required"`
	Notes         string                    `json:"notes"`
	Items         []PrescriptionItemRequest `json:"items" binding:"required"`
}

type PrescriptionItemRequest struct {
	MedicineID   string `json:"medicine_id" binding:"required"`
	Dosage       string `json:"dosage"`
	Frequency    string `json:"frequency"`
	DurationDays int    `json:"duration_days"`
	Quantity     int    `json:"quantity" binding:"required"`
	Instructions string `json:"instructions"`
}

func (s *PrescriptionService) Create(ctx context.Context, req CreatePrescriptionRequest, doctorID string) (*models.PrescriptionWithDetails, error) {
	// Create prescription
	prescription, err := s.repo.Create(ctx, req.AppointmentID, req.DoctorID, req.PatientID, req.Diagnosis, req.Notes)
	if err != nil {
		return nil, errors.New("failed to create prescription")
	}

	// Add items
	var items []models.PrescriptionItem
	for _, item := range req.Items {
		prescriptionItem, err := s.repo.AddItem(ctx, prescription.ID, item.MedicineID, item.Dosage, item.Frequency, item.Instructions, item.DurationDays, item.Quantity)
		if err != nil {
			return nil, errors.New("failed to add prescription item")
		}
		items = append(items, *prescriptionItem)
	}

	return &models.PrescriptionWithDetails{
		Prescription: *prescription,
		Items:        items,
	}, nil
}

func (s *PrescriptionService) GetByID(ctx context.Context, id string) (*models.PrescriptionWithDetails, error) {
	prescription, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("prescription not found")
	}

	items, err := s.repo.GetItems(ctx, id)
	if err != nil {
		return nil, errors.New("failed to get prescription items")
	}

	return &models.PrescriptionWithDetails{
		Prescription: *prescription,
		Items:        items,
	}, nil
}

func (s *PrescriptionService) GetByPatientID(ctx context.Context, patientID string) ([]models.PrescriptionWithDetails, error) {
	return s.repo.GetByPatientID(ctx, patientID)
}

func (s *PrescriptionService) Dispense(ctx context.Context, id string) (*models.Prescription, error) {
	prescription, err := s.repo.Dispense(ctx, id)
	if err != nil {
		return nil, errors.New("failed to dispense prescription")
	}

	// Auto-create billing record
	_ = s.billingRepo.CreateSimple(ctx, prescription.AppointmentID, prescription.PatientID, 500.00)

	return prescription, nil
}
