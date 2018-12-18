package service

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"vincent.com/todo/internal/domain/model"

	"vincent.com/todo/pkg/logger"

	opentracing "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"vincent.com/todo/internal/adapter/http/rpc/helloworld"
	"vincent.com/todo/pkg/auth"
	"vincent.com/todo/pkg/tracing"
)

var log = logger.Logger

//UserRepository -
type UserRepository struct {
	auth *grpc.ClientConn
}

//NewUserRepository -
func NewUserRepository() *UserRepository {
	return &UserRepository{
		auth: auth.InitAuthRPC(tracing.Tracer),
	}
}

//NewToken -
func (r *UserRepository) NewToken(ctx context.Context, u *model.User) (token string, err error) {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "RPC-GetTokenRequest")
	defer span.Finish()
	c := helloworld.NewGreeterClient(r.auth)
	ctx, cancel := context.WithTimeout(childCtx, time.Second)
	defer cancel()
	res, err := c.SayHello(childCtx, &helloworld.HelloRequest{Name: u.GetID()})
	if err != nil {
		log.Sugar().Fatalf("could not greet: %v", err)
		return "", err
	}
	log.Sugar().Infof("Greeting: %s", res.Message)
	return res.Message, nil
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
