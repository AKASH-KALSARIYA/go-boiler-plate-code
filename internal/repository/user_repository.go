package repository

import (
	"go-boilerplate/internal/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetAll() ([]model.User, error)
	Update(user *model.User) error
	Delete(id uint) error
	DeleteAll() error
}

type userRepo struct {
	// db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) Create(user *model.User) error {
	return nil
}

func (r *userRepo) GetByID(id uint) (*model.User, error) {
	return &model.User{}, nil
}

func (r *userRepo) GetAll() ([]model.User, error) {
	return []model.User{}, nil
}

func (r *userRepo) Update(user *model.User) error {
	return nil
}

func (r *userRepo) Delete(id uint) error {
	return nil
}

func (r *userRepo) DeleteAll() error {
	return nil
}
