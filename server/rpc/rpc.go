package rpc

//go:generate protoc -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/googleapis/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:../../ui/src/assets/. rpc.proto
