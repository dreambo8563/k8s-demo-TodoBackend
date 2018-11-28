package auth

import (
	"errors"
	"net/http"
	"os"

	"github.com/imroc/req"
)

// const (
// 	authServiceName = "SERVICE_NAME"
// 	authServicePort = "SERVICE_PORT"
// )

var (
	authServiceName = os.Getenv("SERVICE_NAME")
	authServicePort = os.Getenv("SERVICE_PORT")
)

var authServiceURL = "http://localhost:6000/api/auth/login"

func init() {
	if authServiceName == "" || authServicePort == "" {
		panic("not found auth service config map")
	}
}

// GetToken - get token from auth service
func GetToken(id string) (token string, err error) {
	var reqParam struct {
		ID string `json:"id"`
	}
	reqParam.ID = id

	authServiceURL = "http://" + authServiceName + ":" + authServicePort + "/api/auth/login"
	r, err := req.Post(authServiceURL, req.BodyJSON(&reqParam))
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
