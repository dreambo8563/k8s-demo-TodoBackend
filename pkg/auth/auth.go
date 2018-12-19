package auth

import (
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/resolver"
	"vincent.com/todo/pkg/logger"

	"github.com/imroc/req"
)

var log = logger.Logger
var (
	authServiceURL     = os.Getenv("AUTH_SERVICE_URL")
	authRPCServiceURL  = os.Getenv("AUTH_RPC_SERVICE_URL")
	conn               *grpc.ClientConn
	authServiceBaseURL string
	// authGetTokenURL    string
	authCheckHealthURL string
)

const (
	defautlServiceURL = "localhost:7000"
	defaultRPCURL     = "localhost:50051"
)

//Client -
type Client struct {
	Conn *grpc.ClientConn
}

func init() {
	// set default value
	if authServiceURL == "" {
		authServiceURL = defautlServiceURL
	}
	if authRPCServiceURL == "" {
		authRPCServiceURL = defaultRPCURL
	}
	authServiceBaseURL = "http://" + authServiceURL
	// authGetTokenURL = authServiceBaseURL + "/api/auth/login"
	authCheckHealthURL = authServiceBaseURL + "/healthz"

}

//NewAuthClient -
func NewAuthClient(tracer opentracing.Tracer) *Client {
	if conn != nil && conn.GetState() <= connectivity.Ready {
		return &Client{
			Conn: conn,
		}
	}
	var err error
	log.Info("grpc addr", zap.String("addr", authRPCServiceURL))
	resolver.SetDefaultScheme("dns")
	conn, err = grpc.Dial(authRPCServiceURL, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second), grpc.WithBalancerName(roundrobin.Name), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)), grpc.WithStreamInterceptor(
		otgrpc.OpenTracingStreamClientInterceptor(tracer)))
	if err != nil {
		log.Error("did not connect", zap.String("err", err.Error()))
		return &Client{}
	}
	return &Client{
		Conn: conn,
	}
}

//IsReady -
func (c *Client) IsReady() bool {
	return c.Conn != nil && c.Conn.GetState() == connectivity.Ready
}

// HealthZ - auth service health check
func HealthZ() error {
	_, err := req.Get(authCheckHealthURL)
	return err
}
