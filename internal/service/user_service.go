package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/peace/sandbox/graph/model"
	"github.com/peace/sandbox/internal/mapper"
	"github.com/peace/sandbox/internal/models"
	"github.com/peace/sandbox/internal/repository"
	"github.com/peace/sandbox/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) Users() ([]*model.User, error) {
	users, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	var modelUsers []*model.User

	for _, user := range users {
		modelUsers = append(modelUsers, mapper.MapUserToModel(user))
	}

	return modelUsers, nil
}

func (s *UserService) UserByID(id string) (*model.User, error) {
	user, err := s.Repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserToModel(user), nil
}

func (s *UserService) UserByEmail(email string) (*model.User, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	return mapper.MapUserToModel(user), nil
}

func (s *UserService) Register(input model.RegisterInput) (*model.AuthPayload, error) {
	id := uuid.New()

	role := "USER"

	if input.Role != nil {
		role = string(*input.Role)
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

	payload := &model.AuthPayload{
		Token: token,
		User:  mapper.MapUserToModelWithoutTodos(user),
	}
	return payload, nil
}

func (s *UserService) Login(input model.LoginInput) (*model.AuthPayload, error) {
	user, err := s.Repo.GetByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(user.Password, input.Password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, errors.New("invalid password")
		}

		return nil, err
	}

	token, err := utils.GenerateToken(user.ID.String(), input.Email, user.Role)
	if err != nil {
		return nil, err
	}

	payload := &model.AuthPayload{
		Token: token,
		User:  mapper.MapUserToModelWithoutTodos(user),
	}

	return payload, nil
}
