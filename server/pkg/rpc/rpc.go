package rpc

import (
	google_protobuf "github.com/golang/protobuf/ptypes/any"
	"encoding/json"
	"reflect"
)

//go:generate protoc -I. -I${GOPATH}/src -I${GOPATH}/src/github.com/googleapis/googleapis -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:../../../ui/src/assets/. rpc.proto

func MarshalAny(instance interface{}) (google_protobuf.Any, error) {
	var message google_protobuf.Any

	if instance == nil {
		return message, nil
	}

	value, err := json.Marshal(instance)
	if err != nil {
		return message, err
	}

	message.TypeUrl = reflect.TypeOf(instance).Elem().String()
	message.Value = value

	return message, nil
}
