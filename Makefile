PROJECT_VERSION ?= "0.0.0"

.PHONY: dev
dev:
	go get github.com/googleapis/googleapis
	@which protoc 1> /dev/null || echo "Install protoc: https://github.com/google/protobuf/releases"
	@which dep 1> /dev/null || go get -u github.com/golang/dep/cmd/dep
	@which protoc-gen-grpc-gateway 1> /dev/null || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@which protoc-gen-swagger 1> /dev/null || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@which protoc-gen-go 1> /dev/null || go get -u github.com/golang/protobuf/protoc-gen-go

.PHONY: generate
generate:
	go generate `go list ./server/rpc`

.PHONY: build-server
build-server:
	GO15VENDOREXPERIMENT=1 \
	GOOS=linux \
	go build -ldflags "-X github.com/warmans/kob/cmd/kob-server/main.Version=$(PROJECT_VERSION)" -o build/bin/kob-server ./server/cmd/kob-server
