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
    // Проверяем существует ли пользователь с таким email
    existingUser, err := s.userRepo.FindByEmail(user.Email)
    if err == nil && existingUser != nil {
        return domain.ErrUserAlreadyExists
    }

    // Проверяем существует ли пользователь с таким username
    existingUser, err = s.userRepo.FindByUsername(user.Username)
    if err == nil && existingUser != nil {
        return domain.ErrUsernameAlreadyTaken
    }

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