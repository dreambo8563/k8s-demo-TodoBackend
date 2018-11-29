package models

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"go.uber.org/zap"

	"vincent.com/todo/service/logger"

	"vincent.com/todo/service/auth"
)

var log = logger.Logger

// User struct
type User struct {
	Name  string `json:"name,omitempty"`
	Token string `json:"token,omitempty"`
	ID    string `json:"id,omitempty"`
}

// NewUID - generate uid for a user
func (u *User) NewUID() {
	u.ID = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}

// NewToken - get a token for the user from auth services
func (u *User) NewToken() error {
	if u.ID == "" {
		log.Error("NewToken", zap.String("err", "miss uid"))
		return errors.New("missing user id")
	}
	token, err := auth.GetToken(u.ID)
	if err != nil {
		log.Error("auth.GetToken", zap.String("err", err.Error()))
		return err
	}
	log.Info("get token", zap.String("token", token))
	u.Token = token
	return nil
}
