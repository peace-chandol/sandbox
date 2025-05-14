package service

import (
	"github.com/google/uuid"
	"github.com/peace/sandbox/graph/model"
	"github.com/peace/sandbox/internal/models"
	"github.com/peace/sandbox/internal/repository"
	"github.com/peace/sandbox/utils"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Register(input model.RegisterInput) (*model.AuthPayload, error) {
	id := uuid.New()

	role := "USER"
	var modelRole model.Role = model.RoleUser

	if input.Role != nil {
		role = string(*input.Role)
		modelRole = *input.Role
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     role,
	}

	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}

	token, err := utils.GenerateToken(id.String(), input.Email, role)
	if err != nil {
		return nil, err
	}

	modelUser := &model.User{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  modelRole,
	}

	payload := &model.AuthPayload{
		Token: token,
		User:  modelUser,
	}
	return payload, nil
}
