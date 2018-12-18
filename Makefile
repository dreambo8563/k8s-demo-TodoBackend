v?=latest

local:
	go run cmd/app/main.go
start:
	bash scripts/start.sh
stop:
	bash scripts/stop.sh
docker:
	docker build -f build/Dockerfile -t todo/backend:$(v) .
	docker tag todo/backend:$(v) dreambo8563docker/todo-backend:$(v)
clean:
	rm -f cmd/app/*log app