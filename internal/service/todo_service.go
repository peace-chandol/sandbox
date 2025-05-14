package service

import "github.com/peace/sandbox/internal/repository"

type TodoService struct {
	Repo *repository.TodoRepository
}
