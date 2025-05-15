package mapper

import (
	"github.com/peace/sandbox/graph/model"
	"github.com/peace/sandbox/internal/models"
)

func MapTodoToModel(todo *models.Todo) *model.Todo {
	return &model.Todo{
		ID:   todo.ID.String(),
		Text: todo.Text,
		Done: todo.Done,
	}
}

func MapUserToModel(user *models.User) *model.User {
	var modelTodos []*model.Todo

	for _, todo := range user.Todos {
		modelTodos = append(modelTodos, MapTodoToModel(&todo))
	}

	return &model.User{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  model.Role(user.Role),
		Todo:  modelTodos,
	}
}

func MapUserToModelWithoutTodos(user *models.User) *model.User {
	return &model.User{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
		Role:  model.Role(user.Role),
	}
}
