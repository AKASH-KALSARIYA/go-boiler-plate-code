package service

import (
	"go-boilerplate/internal/dto"
	"go-boilerplate/internal/model"
	"go-boilerplate/internal/repository"
)

type UserService interface {
	Create(req dto.CreateUserRequest) (*dto.UserResponse, error)
	GetByID(id uint) (*dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	Update(id uint, req dto.CreateUserRequest) (*dto.UserResponse, error)
	Delete(id uint) error
	DeleteAll() error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Create(req dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *userService) GetByID(id uint) (*dto.UserResponse, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email}, nil
}

func (s *userService) GetAll() ([]dto.UserResponse, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var res []dto.UserResponse
	for _, u := range users {
		res = append(res, dto.UserResponse{ID: u.ID, Username: u.Username, Email: u.Email})
	}
	return res, nil
}

func (s *userService) Update(id uint, req dto.CreateUserRequest) (*dto.UserResponse, error) {
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
	}
	user.ID = id
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return &dto.UserResponse{ID: user.ID, Username: user.Username, Email: user.Email}, nil
}

func (s *userService) Delete(id uint) error {
	return s.repo.Delete(id)
}

func (s *userService) DeleteAll() error {
	return s.repo.DeleteAll()
}
