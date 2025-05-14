package service

import "github.com/peace/sandbox/internal/repository"

type Service struct {
	Todo *TodoService
	User *UserService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Todo: &TodoService{Repo: repo.Todo},
		User: &UserService{Repo: repo.User},
	}
}
