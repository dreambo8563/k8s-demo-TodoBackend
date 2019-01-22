package service

import (
	"vincent.com/todo/internal/domain/usecase"
	"vincent.com/todo/internal/pkg/logger"
)

var log = logger.Logger()

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
