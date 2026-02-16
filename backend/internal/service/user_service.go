package service

import (
	"errors"

	"github.com/hadi-projects/go-react-starter/config"
	"github.com/hadi-projects/go-react-starter/internal/dto"
	"github.com/hadi-projects/go-react-starter/internal/entity"
	"github.com/hadi-projects/go-react-starter/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dto.RegisterRequest) (*dto.UserResponse, error)
	GetMe(userID uint) (*dto.UserResponse, error)
}

type userService struct {
	userRepo repository.UserRepository
	config   *config.Config
}

func NewUserService(userRepo repository.UserRepository, config *config.Config) UserService {
	return &userService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *userService) Register(req dto.RegisterRequest) (*dto.UserResponse, error) {
	// Check if email exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.config.Security.BCryptCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		RoleID:   2, // Default to user role
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *userService) GetMe(userID uint) (*dto.UserResponse, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
