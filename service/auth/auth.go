package auth

import (
	"errors"
	"net/http"
	"os"

	"vincent.com/todo/service/logger"

	"github.com/imroc/req"
)

var log = logger.Logger
var (
	authServiceName = os.Getenv("SERVICE_NAME")
	authServicePort = os.Getenv("SERVICE_PORT")

	authServiceBaseURL string
	authGetTokenURL    string
	authCheckHealthURL string
)

func init() {
	// set default value
	if authServiceName == "" {
		authServiceName = "localhost"
	}
	if authServicePort == "" {
		authServicePort = "6000"
	}
	authServiceBaseURL = "http://" + authServiceName + ":" + authServicePort
	authGetTokenURL = authServiceBaseURL + "/api/auth/login"
	authCheckHealthURL = authServiceBaseURL + "/healthz"
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

// HealthZ - auth service health check
func HealthZ() error {
	_, err := req.Get(authCheckHealthURL)
	return err
}
