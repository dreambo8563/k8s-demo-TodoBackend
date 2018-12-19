package service

import (
	"vincent.com/todo/internal/domain/usecase"
)

//InitializeUserCase -
func InitializeUserCase() *usecase.UserUsecase {
	repo := NewUserRepository()
	return usecase.NewUserUsecase(repo)
}
