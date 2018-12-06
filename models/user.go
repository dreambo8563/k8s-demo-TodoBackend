package models

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"

	opentracing "github.com/opentracing/opentracing-go"
	"vincent.com/todo/service/auth"
	"vincent.com/todo/service/logger"
)

var log = logger.Logger

// User struct
type User struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
	ID    string `json:"id,omitempty"`
}

// NewUID - generate uid for a user
func (u *User) NewUID(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "NewUID")
	defer span.Finish()
	u.ID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}

// NewToken - get a token for the user from auth services
func (u *User) NewToken(ctx context.Context) error {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "GetNewToken")
	defer span.Finish()
	if u.ID == "" {
		span.LogKV("event", "verify", "err", "miss uid")
		log.Error("NewToken", zap.String("err", "miss uid"))
		return errors.New("missing user id")
	}
	token, err := auth.GetToken(childCtx, u.ID)
	if err != nil {
		log.Error("auth.GetToken", zap.String("err", err.Error()))
		return err
	}
	log.Info("get token", zap.String("token", token))
	u.Token = token
	return nil
}

// RPCNewToken - get a token for the user from auth services
func (u *User) RPCNewToken(ctx context.Context) error {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "GetNewToken")
	defer span.Finish()
	if u.ID == "" {
		span.LogKV("event", "verify", "err", "miss uid")
		log.Error("NewToken", zap.String("err", "miss uid"))
		return errors.New("missing user id")
	}
	token, err := auth.RPCGetToken(childCtx, u.ID)
	if err != nil {
		log.Error("auth.GetToken", zap.String("err", err.Error()))
		return err
	}
	log.Info("get token", zap.String("token", token))
	u.Token = token
	return nil
}
