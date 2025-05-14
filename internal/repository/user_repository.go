package repository

import (
	"github.com/peace/sandbox/graph/model"
	"github.com/peace/sandbox/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) GetById(id string) ([]*model.User, error) {
	return nil, nil
}

func (r *UserRepository) GetByEmail(email string) ([]*model.User, error) {
	return nil, nil
}

func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
