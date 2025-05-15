package repository

import (
	"github.com/peace/sandbox/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.DB.Preload("Todos").Find(&users).Error
	return users, err
}

func (r *UserRepository) GetById(id string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Todos").First(&user, "id = ?", id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Preload("Todos").First(&user, "email = ?", email).Error
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) Update(user *models.User) error {
	return r.DB.Save(user).Error
}

func (r *UserRepository) Delete(id string) error {
	return r.DB.Delete(&models.User{}, "id = ?", id).Error
}
