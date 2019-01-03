package service

import (
	"context"

	"github.com/rs/xid"
)

//UUIDRepository -
type UUIDRepository struct {
}

//NewUUIDRepository -
func NewUUIDRepository() *UUIDRepository {
	return &UUIDRepository{}
}

//New -
func (r *UUIDRepository) New(ctx context.Context) string {
	guid := xid.New()
	log.Sugar().Infof("machine id: %v,", guid.Machine())
	return guid.String()
}
