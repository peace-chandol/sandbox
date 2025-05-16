package repository

import (
	"github.com/peace/sandbox/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
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

func (r *UserRepository) Update(userInput *models.User) (*models.User, error) {
	var user models.User
	if result := r.DB.Where("id = ?", userInput.ID).First(&user); result.Error != nil {
		return nil, result.Error
	}

	user.Name = userInput.Name
	user.Password = userInput.Password

	if err := r.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Todo{}, "user_id = ?", id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.User{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}
