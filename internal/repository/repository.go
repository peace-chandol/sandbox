package repository

import "gorm.io/gorm"

type Repository struct {
	Todo *TodoRepository
	User *UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Todo: &TodoRepository{DB: db},
		User: &UserRepository{DB: db},
	}
}
