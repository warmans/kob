PROJECT_VERSION ?= "0.0.0"

DEV_DB_USER ?= kob
DEV_DB_PASS ?= insecure
DEV_DB_HOST ?= localhost
DEV_DB_DSN ?= "postgres://$(DEV_DB_USER):$(DEV_DB_PASS)@$(DEV_DB_HOST)/$(DEV_DB_USER)?sslmode=disable"

.PHONY: dev
dev:
	go get github.com/googleapis/googleapis
	@which protoc 1> /dev/null || echo "Install protoc: https://github.com/google/protobuf/releases"
	@which dep 1> /dev/null || go get -u github.com/golang/dep/cmd/dep
	@which protoc-gen-grpc-gateway 1> /dev/null || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@which protoc-gen-swagger 1> /dev/null || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@which protoc-gen-go 1> /dev/null || go get -u github.com/golang/protobuf/protoc-gen-go
	@which goose 1> /dev/null || go get -u github.com/pressly/goose/cmd/goose

.PHONY: generate
generate:
	@go generate `go list ./server/pkg/rpc`
	@cd ui; node_modules/.bin/a4apigen -s src/assets/rpc.swagger.json -o src/app/service/api-client

.PHONY: build-server
build-server:
	GO15VENDOREXPERIMENT=1 \
	GOOS=linux \
	go build -ldflags "-X github.com/warmans/kob/cmd/kob-server/main.Version=$(PROJECT_VERSION)" -o build/bin/kob-server ./server/cmd/kob-server

.PHONY: run
run:
	@$(MAKE) -j2 run-ui run-server

.PHONY: run-ui
run-ui:
	cd ui; node_modules/.bin/ng serve -pc proxy.dev.conf.json

.PHONY: run-server
run-server:
	./build/bin/kob-server -db-dsn=$(DEV_DB_DSN) -db-query-log-enabled=true

.PHONY: run-services
run-services: 
	cd dev; POSTGRES_USER=$(DEV_DB_USER) POSTGRES_PASSWORD=$(DEV_DB_PASS) docker-compose up

.PHONY: migrate-up
migrate-up: 
	goose -dir migrations postgres $(DEV_DB_DSN) up

.PHONY: migrate-down
migrate-down: 
	goose -dir migrations postgres $(DEV_DB_DSN) down

.PHONY: generate_private_key
generate-private-key:
	@mkdir -p var/keys
	@ssh-keygen -t rsa -b 4096 -f var/keys/jwtRS256.key

.PHONY: dev-dsn
dev-dsn:
	@echo $(DEV_DB_DSN)
