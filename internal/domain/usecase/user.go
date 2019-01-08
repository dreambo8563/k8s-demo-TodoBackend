package usecase

import (
	"context"

	"vincent.com/todo/internal/domain/model"

	"vincent.com/todo/internal/domain/repository"
)

// User -
type User struct {
	ID   string `json:"uid"`
	Name string `json:"username,omitempty"`
}

//IUserUsecase -
type IUserUsecase interface {
	RegisterUser(ctx context.Context, name, password string) (token string, err error)
	GetInfo(ctx context.Context, token string) *User
}

//UserUsecase -
type UserUsecase struct {
	repo repository.UserRepository
}

//NewUserUsecase -
func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

//RegisterUser -
func (u *UserUsecase) RegisterUser(ctx context.Context, name, password string) (user *User, token string, err error) {
	var userItem *model.User
	userItem, err = u.repo.CreateUser(ctx, name, password)
	if err != nil {
		return nil, "", err
	}
	token, err = u.repo.NewToken(ctx, userItem)
	if err != nil {
		return nil, "", err
	}

	return toUser(userItem), token, nil
}

//GetInfo -
func (u *UserUsecase) GetInfo(ctx context.Context, token string) (*User, error) {
	user, err := u.repo.ParseToken(ctx, token)
	if err != nil {
		return nil, err
	}
	// get info by user.ID
	return toUser(user), nil
}

func toUser(user *model.User) *User {
	return &User{
		ID:   user.GetID(),
		Name: user.GetName(),
	}
}
