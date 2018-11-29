package auth

import (
	"errors"
	"net/http"
	"os"

	"github.com/imroc/req"
)

var (
	authServiceName = os.Getenv("SERVICE_NAME")
	authServicePort = os.Getenv("SERVICE_PORT")

	authServiceBaseURL = "http://" + authServiceName + ":" + authServicePort
	authGetTokenURL    = authServiceBaseURL + "/api/auth/login"
)

func init() {
	if authServiceName == "" {
		authServiceName = "localhost"
	}
	if authServicePort == "" {
		authServicePort = "6000"
	}
}

// GetToken - get token from auth service
func GetToken(id string) (token string, err error) {
	var reqParam struct {
		ID string `json:"id"`
	}
	reqParam.ID = id

	r, err := req.Post(authGetTokenURL, req.BodyJSON(&reqParam))
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
