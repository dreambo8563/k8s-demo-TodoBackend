v?=latest

local:
	export JAEGER_SERVICE_NAME=aaa && go run cmd/app/main.go
start:
	bash scripts/start.sh
stop:
	bash scripts/stop.sh
docker:
	docker build -f build/Dockerfile -t todo/backend:$(v) .
	docker tag todo/backend:$(v) dreambo8563docker/todo-backend:$(v)
	docker push dreambo8563docker/todo-backend:$(v)
rpc:
	protoc -I internal/adapter/http/rpc/auth/ internal/adapter/http/rpc/auth/auth.proto --go_out=plugins=grpc:internal/adapter/http/rpc/auth
clean:
	rm -f todobackend-log cmd/app/*log app
	rm -f todobackend-log internal/pkg/uc/*.db