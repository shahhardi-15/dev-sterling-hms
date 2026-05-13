package services

import (
	"context"
	"errors"
	"sterling-hms/internal/models"
	"sterling-hms/internal/repositories"
)

type AppointmentService struct {
	repo *repositories.AppointmentRepository
}

func NewAppointmentService(repo *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{repo: repo}
}

type CreateAppointmentRequest struct {
	PatientID   string `json:"patient_id" binding:"required"`
	DoctorID    string `json:"doctor_id" binding:"required"`
	ScheduledAt string `json:"scheduled_at" binding:"required"`
	Type        string `json:"type" binding:"required,oneof=online walk_in follow_up emergency"`
	Reason      string `json:"reason"`
}

func (s *AppointmentService) GetAll(ctx context.Context) ([]models.AppointmentWithDetails, error) {
	return s.repo.GetAll(ctx)
}

func (s *AppointmentService) GetByID(ctx context.Context, id string) (*models.AppointmentWithDetails, error) {
	appointment, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.New("appointment not found")
	}
	return appointment, nil
}

func (s *AppointmentService) GetByPatientID(ctx context.Context, patientID string) ([]models.AppointmentWithDetails, error) {
	return s.repo.GetByPatientID(ctx, patientID)
}

func (s *AppointmentService) GetByDoctorID(ctx context.Context, doctorID string) ([]models.AppointmentWithDetails, error) {
	return s.repo.GetByDoctorID(ctx, doctorID)
}

func (s *AppointmentService) Create(ctx context.Context, req CreateAppointmentRequest, createdBy string) (*models.Appointment, error) {
	appointment, err := s.repo.Create(ctx, req.PatientID, req.DoctorID, req.ScheduledAt, req.Type, req.Reason, createdBy)
	if err != nil {
		return nil, errors.New("failed to create appointment")
	}
	return appointment, nil
}

func (s *AppointmentService) Approve(ctx context.Context, id string) (*models.Appointment, error) {
	appointment, err := s.repo.UpdateStatus(ctx, id, "approved")
	if err != nil {
		return nil, errors.New("failed to approve appointment")
	}
	return appointment, nil
}

func (s *AppointmentService) Reject(ctx context.Context, id string) (*models.Appointment, error) {
	appointment, err := s.repo.UpdateStatus(ctx, id, "rejected")
	if err != nil {
		return nil, errors.New("failed to reject appointment")
	}
	return appointment, nil
}

func (s *AppointmentService) Cancel(ctx context.Context, id string) (*models.Appointment, error) {
	appointment, err := s.repo.UpdateStatus(ctx, id, "cancelled")
	if err != nil {
		return nil, errors.New("failed to cancel appointment")
	}
	return appointment, nil
}

func (s *AppointmentService) Complete(ctx context.Context, id string) (*models.Appointment, error) {
	appointment, err := s.repo.UpdateStatus(ctx, id, "completed")
	if err != nil {
		return nil, errors.New("failed to complete appointment")
	}
	return appointment, nil
}
