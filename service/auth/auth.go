package auth

import (
	"errors"
	"net/http"

	"github.com/imroc/req"
)

// GetToken - get token from auth service
func GetToken(id string) (token string, err error) {
	var reqParam struct {
		ID string `json:"id"`
	}
	reqParam.ID = id
	r, err := req.Post("http://todo-auth-service:6000/api/auth/login", req.BodyJSON(&reqParam))
	if err != nil {
		return "", err
	}
	if r.Response().StatusCode != http.StatusOK {
		var msg struct {
			Msg string `json:"msg"`
		}
		err = r.ToJSON(msg)
		return "", errors.New(msg.Msg)
	}

	var resParam struct {
		Token string `json:"token"`
	}

	err = r.ToJSON(&resParam)
	if err != nil {
		return "", err
	}
	return resParam.Token, nil
}
