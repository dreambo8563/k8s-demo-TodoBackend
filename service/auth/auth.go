package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"

	"vincent.com/todo/rpc/helloworld"

	"google.golang.org/grpc"

	"vincent.com/todo/service/logger"

	"github.com/imroc/req"
)

var log = logger.Logger
var (
	authServiceURL     = os.Getenv("AUTH_SERVICE_URL")
	authRPCServiceURL  = os.Getenv("AUTH_RPC_SERVICE_URL")
	conn               *grpc.ClientConn
	authServiceBaseURL string
	authGetTokenURL    string
	authCheckHealthURL string
)

const (
	name              = "world"
	defautlServiceURL = "localhost:7000"
	defaultRPCURL     = "localhost:50051"
)

func init() {
	// set default value
	if authServiceURL == "" {
		authServiceURL = defautlServiceURL
	}
	if authRPCServiceURL == "" {
		authRPCServiceURL = defaultRPCURL
	}
	authServiceBaseURL = "http://" + authServiceURL
	authGetTokenURL = authServiceBaseURL + "/api/auth/login"
	authCheckHealthURL = authServiceBaseURL + "/healthz"

}

//InitAuthRPC -
func InitAuthRPC() *grpc.ClientConn {
	var err error
	log.Info("grpc addr", zap.String("addr", authRPCServiceURL))
	conn, err = grpc.Dial(authRPCServiceURL, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err != nil {
		log.Fatal("did not connect", zap.String("err", err.Error()))
	}
	return conn
}

// GetToken - get token from auth service
func GetToken(id string) (token string, err error) {
	var reqParam struct {
		ID string `json:"id"`
	}
	reqParam.ID = id
	hello()
	log.Info("http addr", zap.String("addr", authServiceBaseURL))
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

func hello() {
	c := helloworld.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Sugar().Fatalf("could not greet: %v", err)
	}
	log.Sugar().Infof("Greeting: %s", r.Message)
}
