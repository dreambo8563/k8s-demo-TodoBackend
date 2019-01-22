package repository

import (
	"context"

	"vincent.com/todo/internal/domain/model"
)

//UserRepository - User repo interface
type UserRepository interface {
	NewToken(context.Context, *model.User) (token string, err error)
	CreateUser(context.Context, string, string) (*model.User, error)
	ParseToken(context.Context, string) (*model.User, error)
	GetUser(context.Context, string) (*model.User, error)
	IsDup(context.Context, string) (bool, error)
}
