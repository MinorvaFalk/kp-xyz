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
	goose -dir migrations status

migrate-up:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migrations up

migrate-down:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migrations down

migrate-reset:
	export GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING); \
	goose -dir migrations reset

run:
	go run cmd/api/*.go

.PHONY: mocks
mocks:
	sleep 1 && rm -rfd mocks && mockery

build:
	docker build -t ${IMAGE_NAME}:${IMAGE_VERSION} .