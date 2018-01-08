package server

import (
	"github.com/warmans/kob/server/pkg/search"
	"github.com/warmans/kob/server/pkg/auth/token"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/warmans/kob/server/pkg/db"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/warmans/kob/server/pkg/rpc/server/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
)

func NewRPCServer(
	logger *zap.Logger, 
	store *db.Store, 
	tokenStrategy token.Strategy,
	index *search.Index,
) *Server {
	
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_zap.StreamServerInterceptor(logger),
			//grpc_recovery.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(token.AuthFunc(tokenStrategy, logger)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_zap.UnaryServerInterceptor(logger),
			//grpc_recovery.UnaryServerInterceptor(),
			grpc_auth.UnaryServerInterceptor(token.AuthFunc(tokenStrategy, logger)),
		)),
	)

	grpc_prometheus.Register(grpcServer)

	rpc.RegisterKobServiceServer(grpcServer, service.NewKobService(store, index))
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
