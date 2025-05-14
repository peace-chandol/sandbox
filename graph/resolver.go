package graph

import (
	"github.com/peace/sandbox/internal/repository"
	"github.com/peace/sandbox/internal/service"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service *service.Service
}

func NewResolver(db *gorm.DB) *Resolver {
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	return &Resolver{Service: svc}
}
