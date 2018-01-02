package service

import (
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NotFound() error {
	return status.Error(codes.NotFound, "resource was not found")
}

func Err2Code(err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return NotFound()
	}
	return status.Error(codes.Internal, err.Error())
}
