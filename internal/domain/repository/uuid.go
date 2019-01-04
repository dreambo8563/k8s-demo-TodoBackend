package repository

import (
	"context"
)

//UUIDRepository - U
type UUIDRepository interface {
	New(context.Context) (uuid string)
}
