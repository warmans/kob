package server

import (
	"net"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/warmans/kob/server/pkg/rpc/server/service"
)

func NewRPCServer(logger *zap.Logger) *Server {

	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	grpc_prometheus.Register(grpcServer)

	rpc.RegisterKobServiceServer(grpcServer, service.NewKobService())

	return &Server{srv: grpcServer}
}

type Server struct {
	srv *grpc.Server
	log *zap.Logger
}

func (s *Server) Serve(bindAddr string, bindPort int) error {
	
	bind := fmt.Sprintf("%s:%d", bindAddr, bindPort)
	conn, err := net.Listen("tcp", bind)
	if err != nil {
		return err
	}
	return s.srv.Serve(conn)
}
