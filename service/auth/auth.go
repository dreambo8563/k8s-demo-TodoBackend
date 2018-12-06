package auth

import (
	"context"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"vincent.com/todo/rpc/helloworld"
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
func InitAuthRPC(tracer opentracing.Tracer) *grpc.ClientConn {
	var err error
	log.Info("grpc addr", zap.String("addr", authRPCServiceURL))
	conn, err = grpc.Dial(authRPCServiceURL, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)), grpc.WithStreamInterceptor(
		otgrpc.OpenTracingStreamClientInterceptor(tracer)))
	if err != nil {
		log.Fatal("did not connect", zap.String("err", err.Error()))
	}
	return conn
}

// GetToken - get token from auth service
func GetToken(ctx context.Context, id string) (token string, err error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "HTTP-GetTokenRequest")
	defer span.Finish()
	var reqParam struct {
		ID string `json:"id"`
	}
	reqParam.ID = id
	header := make(http.Header)
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, authGetTokenURL)
	ext.HTTPMethod.Set(span, "POST")
	span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header),
	)

	log.Info("http addr", zap.String("addr", authServiceBaseURL))
	r, err := req.Post(authGetTokenURL, header, req.BodyJSON(&reqParam))
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

func hello(ctx context.Context, id string) (token string, err error) {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "helloRPC")
	defer span.Finish()
	c := helloworld.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(childCtx, time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: id})
	if err != nil {
		log.Sugar().Fatalf("could not greet: %v", err)
		return "", err
	}
	log.Sugar().Infof("Greeting: %s", r.Message)
	return r.Message, nil
}

// RPCGetToken - get token from auth service using RPC
func RPCGetToken(ctx context.Context, id string) (token string, err error) {
	span, childCtx := opentracing.StartSpanFromContext(ctx, "RPC-GetTokenRequest")
	defer span.Finish()
	return hello(childCtx, id)
}
