GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= admin:admin@/database?parseTime=true

IMAGE_NAME ?= kp/xyz
IMAGE_VERSION ?= 0.0.0

mysql:
	docker run -itd \
	--name mysql \
	-e MYSQL_ROOT_PASSWORD=root \
	-e MYSQL_USER=admin \
	-e MYSQL_PASSWORD=admin \
	-e MYSQL_DATABASE=database \
	-e TZ=Asia/Jakarta \
	-p 3306:3306 \
	mysql:5.7

migrate-status:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migration status

migrate-up:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migration up

migrate-down:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migration down

migrate-reset:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migration reset

run:
	go run cmd/api/*.go

swag:
	swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/api/main.go

build:
	docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .