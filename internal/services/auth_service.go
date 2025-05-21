package services

import (
	"errors"
	"todo-app/internal/domain"
	"todo-app/internal/repository"
	"todo-app/pkg/jwt"
)

type AuthService interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtUtil  jwt.JWTUtil
}

func NewAuthService(userRepo repository.UserRepository, jwtUtil jwt.JWTUtil) AuthService {
	return &authService{userRepo: userRepo, jwtUtil: jwtUtil}
}

func (s *authService) Register(user *domain.User) error {
	return s.userRepo.Create(user)
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email")
	}

	if user.Password != password {
		return "", errors.New("invalid password")
	}

	token, err := s.jwtUtil.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}