package usecase

import (
	"errors"
	"golang-clean-arch-api/internal/entity"
	"golang-clean-arch-api/internal/repository"
)

// Service is an interface from which our api module can access our repository of all our models
type UserUseCase interface {
	GetAll() ([]*entity.User, error)
	GetByID(id string) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(id string) error
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: r,
	}
}

func (u *userUseCase) GetAll() ([]*entity.User, error) {
	return u.userRepo.GetAll()
}

func (u *userUseCase) GetByID(id string) (*entity.User, error) {
	return u.userRepo.GetByID(id)
}

func (u *userUseCase) Create(user *entity.User) error {
	if user == nil {
		return errors.New("user can't be empty")
	}

	return u.userRepo.Create(user)
}

func (u *userUseCase) Update(user *entity.User) error {
	if user == nil {
		return errors.New("user can't be empty")
	}

	return u.userRepo.Update(user)
}

func (u *userUseCase) Delete(id string) error {
	return u.userRepo.Delete(id)
}
