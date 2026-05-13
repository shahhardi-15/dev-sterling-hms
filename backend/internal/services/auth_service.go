package services

import (
	"context"
	"errors"
	"sterling-hms/internal/repositories"
	"sterling-hms/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         interface{} `json:"user"`
	ExpiresIn    int         `json:"expires_in"`
}

func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// Find user by email
	user, err := s.repo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	accessToken, err := utils.GenerateToken(user.ID, user.Role, user.FullName)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: map[string]interface{}{
			"id":        user.ID,
			"full_name": user.FullName,
			"email":     user.Email,
			"role":      user.Role,
		},
		ExpiresIn: 3600,
	}, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	user, err := s.repo.FindUserByID(ctx, userID)
	if err != nil {
		return errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword))
	if err != nil {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	return s.repo.UpdatePassword(ctx, userID, string(hashedPassword))
}
