package repositories

import "github.com/MarcKVR/clean_arquitecture/domain/models"

type UserRepository interface {
	Create(user *models.User) error
	FindById(id uint) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}
