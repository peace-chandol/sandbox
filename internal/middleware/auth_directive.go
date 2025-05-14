package middleware

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/peace/sandbox/graph/model"
)

func AuthDirective(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (interface{}, error) {
	user := ForContext(ctx)
	if user == nil {
		return nil, errors.New("Unauthenticated: user not found")
	}

	userRole := model.Role(user.Role)
	for _, r := range roles {
		if userRole == r {
			return next(ctx)
		}
	}

	return nil, fmt.Errorf("Forbidden: required role %v but user has role %s", roles, user.Role)
}
