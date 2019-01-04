package service

import (
	"vincent.com/todo/internal/domain/usecase"
)

//InitializeUserCase -
func InitializeUserCase() *usecase.UserUsecase {
	repo := NewUserRepository()
	return usecase.NewUserUsecase(repo)
}

//InitializeUUIDCase -
func InitializeUUIDCase() *usecase.UUIDUsecase {
	repo := NewUUIDRepository()
	return usecase.NewUUIDUsecase(repo)
}
