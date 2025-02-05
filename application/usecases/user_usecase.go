package usecases

import (
	"github.com/MarcKVR/clean_arquitecture/domain/models"
	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
)

type UserUseCase struct {
	repo repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) CreateUser(user *models.User) error {
	return u.repo.Create(user)
}

func (u *UserUseCase) GetUserById(id string) (*models.User, error) {
	return u.repo.FindById(id)
}

func (u *UserUseCase) GetAllUsers() ([]models.User, error) {
	return u.repo.FindAll()
}

func (u *UserUseCase) UpdateUser(user *models.User) error {
	return u.repo.Update(user)
}

func (u *UserUseCase) DeleteUser(id uint) error {
	return u.repo.Delete(id)
}
