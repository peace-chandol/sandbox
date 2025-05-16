package repository

import "gorm.io/gorm"

type TodoRepository struct {
	DB *gorm.DB
}
