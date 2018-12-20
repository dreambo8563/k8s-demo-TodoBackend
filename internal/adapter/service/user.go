package service

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"vincent.com/todo/internal/domain/model"

	"vincent.com/todo/internal/pkg/logger"

	opentracing "github.com/opentracing/opentracing-go"
	authService "vincent.com/todo/internal/adapter/http/rpc/auth"
	"vincent.com/todo/internal/pkg/auth"
	"vincent.com/todo/internal/pkg/tracing"
)

var log = logger.Logger()

//UserRepository -
type UserRepository struct {
	auth *auth.Client
}

//NewUserRepository -
func NewUserRepository() *UserRepository {
	return &UserRepository{
		auth: auth.NewAuthClient(tracing.NewTraceClient().Tracer),
	}
}

//NewToken -
func (r *UserRepository) NewToken(ctx context.Context, u *model.User) (token string, err error) {
	if !r.auth.IsReady() {
		return "", errors.New("can not connect to auth RPC server")
	}
	span, childCtx := tracing.StartSpanFromContext(ctx, "RPC-GetTokenRequest")
	defer span.Finish()
	c := authService.NewAuthServiceClient(r.auth.Conn)
	ctx, cancel := context.WithTimeout(childCtx, time.Second)
	defer cancel()
	res, err := c.GetToken(childCtx, &authService.GetTokenRequest{Uid: u.GetID()})
	if err != nil {
		log.Sugar().Fatalf("could not greet: %v", err)
		return "", err
	}
	log.Sugar().Infof("Greeting: %s", res.Token)
	return res.Token, nil
}

//CreateUser -
func (r *UserRepository) CreateUser(ctx context.Context, name, password string) (*model.User, error) {
	id := newUID(ctx)
	return &model.User{
		Name:     name,
		Password: password,
		ID:       id,
	}, nil
}

// generate uid for a user / fake
func newUID(ctx context.Context) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewUID")
	defer span.Finish()
	return strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
