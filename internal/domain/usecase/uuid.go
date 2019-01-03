package usecase

import (
	"context"

	"vincent.com/todo/internal/domain/repository"
)

//UUIDUsecase -
type UUIDUsecase struct {
	repo repository.UUIDRepository
}

//NewUUIDUsecase -
func NewUUIDUsecase(repo repository.UUIDRepository) *UUIDUsecase {
	return &UUIDUsecase{
		repo: repo,
	}
}

//New -
func (r *UUIDUsecase) New(ctx context.Context) string {
	return r.repo.New(ctx)
}
