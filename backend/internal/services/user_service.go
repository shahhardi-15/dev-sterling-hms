package services

import (
	"context"
	"errors"
	"sterling-hms/internal/models"
	"sterling-hms/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

type CreateUserRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=admin doctor patient receptionist pharmacist"`
}

type UpdateUserRequest struct {
        FullName string `json:"full_name" binding:"required"`
        IsActive bool   `json:"is_active"`
        Password string `json:"password"`
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user, err := s.repo.CreateUser(ctx, req.FullName, req.Email, string(hashedPassword), req.Role)
	if err != nil {
		return nil, errors.New("failed to create user — email may already exist")
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, req UpdateUserRequest) (*models.User, error) {
        if req.Password != "" {
                hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
                if err != nil {
                        return nil, errors.New("failed to hash password")
                }
                if err := s.repo.UpdatePassword(ctx, id, string(hashedPassword)); err != nil {
                        return nil, errors.New("failed to update password")
                }
        }
        user, err := s.repo.UpdateUser(ctx, id, req.FullName, req.IsActive)
        if err != nil {
                return nil, errors.New("failed to update user")
        }
        return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) GetAllPatients(ctx context.Context) ([]models.PatientWithUser, error) {
	return s.repo.GetAllPatients(ctx)
}

func (s *UserService) GetAllDoctors(ctx context.Context) ([]models.DoctorWithUser, error) {
	return s.repo.GetAllDoctors(ctx)
}

func (s *UserService) GetDoctorByUserID(ctx context.Context, userID string) (*models.DoctorWithUser, error) {
	doctor, err := s.repo.GetDoctorByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("doctor profile not found")
	}
	return doctor, nil
}
