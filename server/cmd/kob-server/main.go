package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/warmans/kob/server/rpc"
	"github.com/warmans/kob/server/rpc/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	Version = "unknown"
)

var (
	bindAddr = flag.String("bind-addr", "", "bind to this interface")
	httpPort = flag.Int("http-port", 8080, "port to bind HTTP server")
	grpcPort = flag.Int("grcp-port", 9090, "port to bind gRPC server")
	webRoot  = flag.String("web-root", "build", "path to public dir")
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//start GRPC
	go serveGRPC(logger)

	// start HTTP
	serveHTTP(logger)
}

func serveGRPC(logger *zap.Logger) {
	grpcServer := grpc.NewServer()
	rpc.RegisterKobServiceServer(grpcServer, service.NewKobService())

	bind := fmt.Sprintf("%s:%d", *bindAddr, *grpcPort)
	conn, err := net.Listen("tcp", bind)
	if err != nil {
		logger.Fatal("gRPC Failed", zap.Error(err))
	}

	logger.Info("Starting GRPC...", zap.String("bind", bind))
	logger.Fatal("gRPC Failed", zap.Error(grpcServer.Serve(conn)))
}

func serveHTTP(logger *zap.Logger) {

	grpcConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", *grpcPort), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial GRPC Server", zap.Error(err))
	}
	defer grpcConn.Close()

	gwmux := runtime.NewServeMux()
	rpc.RegisterKobServiceHandler(context.Background(), gwmux, grpcConn)

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	bind := fmt.Sprintf("%s:%d", *bindAddr, *httpPort)

	logger.Info("Starting HTTP...", zap.String("bind", bind))
	logger.Fatal("HTTP Failed", zap.Error(http.ListenAndServe(bind, mux)))
}
