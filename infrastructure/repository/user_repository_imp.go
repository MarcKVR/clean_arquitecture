package repository

import (
	"github.com/MarcKVR/clean_arquitecture/domain/models"
	"github.com/MarcKVR/clean_arquitecture/domain/repositories"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	gorm *gorm.DB
}

func NewUserRepository(gorm *gorm.DB) repositories.UserRepository {
	return &UserRepositoryGorm{gorm: gorm}
}

func (r *UserRepositoryGorm) Create(user *models.User) error {
	return r.gorm.Create(user).Error
}

func (r *UserRepositoryGorm) FindById(id uint) (*models.User, error) {
	var user models.User
	resul := r.gorm.First(&user, id)

	return &user, resul.Error
}

func (r *UserRepositoryGorm) FindAll() ([]models.User, error) {
	var users []models.User
	resul := r.gorm.Find(&users)

	return users, resul.Error
}

func (r *UserRepositoryGorm) Update(user *models.User) error {
	return r.gorm.Save(user).Error
}

func (r *UserRepositoryGorm) Delete(id uint) error {
	return r.gorm.Delete(&models.User{}, id).Error
}
