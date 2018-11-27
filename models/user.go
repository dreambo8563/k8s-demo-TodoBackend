package models

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"vincent.com/todo/service/auth"
)

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
		return errors.New("missing user id")
	}
	token, err := auth.GetToken(u.ID)
	if err != nil {
		return err
	}
	u.Token = token
	return nil
}
