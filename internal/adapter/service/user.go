package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"vincent.com/todo/internal/pkg/uc"

	"vincent.com/todo/internal/domain/model"

	opentracing "github.com/opentracing/opentracing-go"
	authService "vincent.com/todo/internal/adapter/http/rpc/auth"
	"vincent.com/todo/internal/pkg/auth"
	"vincent.com/todo/internal/pkg/tracing"
)

//UserRepository -
type UserRepository struct {
	auth *auth.Client
	db   *uc.Client
}

//NewUserRepository -
func NewUserRepository() *UserRepository {
	return &UserRepository{
		auth: auth.NewAuthClient(tracing.NewTraceClient().Tracer),
		db:   uc.NewDB(),
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
	res, err := c.GetToken(ctx, &authService.GetTokenRequest{Uid: u.GetID()})
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
	fmt.Println("id", id)
	fmt.Println("db", r.db)
	err := r.db.Save(&model.User{
		Name:     name,
		Password: password,
		ID:       id,
	})
	fmt.Println("CreateUser", err)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Name:     name,
		Password: password,
		ID:       id,
	}, nil
}

//ParseToken -
func (r *UserRepository) ParseToken(ctx context.Context, token string) (*model.User, error) {
	if !r.auth.IsReady() {
		return nil, errors.New("can not connect to auth RPC server")
	}
	c := authService.NewAuthServiceClient(r.auth.Conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	res, err := c.ParseToken(ctx, &authService.ParseTokenRequest{
		Token: token,
	})
	if err != nil {
		return nil, err
	}
	log.Info("parse token get id", log.String("id", res.Id))
	return &model.User{
		ID: res.Id,
	}, nil
}

//GetUser -
func (r *UserRepository) GetUser(ctx context.Context, uid string) (*model.User, error) {
	user := &model.User{
		ID: uid,
	}
	fmt.Println(uid)
	err := r.db.Get(user)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// generate uid for a user / fake
func newUID(ctx context.Context) string {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewUID")
	defer span.Finish()
	return strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
