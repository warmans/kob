package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"

	"github.com/gocraft/dbr"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/warmans/kob/server/pkg/db"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/warmans/kob/server/pkg/rpc/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "github.com/lib/pq"
)

var (
	Version = "unknown"
)

var (
	bindAddr = flag.String("bind-addr", "", "bind to this interface")
	httpPort = flag.Int("http-port", 8080, "port to bind HTTP server")
	grpcPort = flag.Int("grcp-port", 9090, "port to bind gRPC server")
	webRoot  = flag.String("web-root", "build/dist", "path to public dir")
	dbDSN    = flag.String("db-dsn", "", "DSN of postgres DB")
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn, err := dbr.Open("postgres", *dbDSN, nil)
	if err != nil {
		logger.Fatal("DB connection not possible", zap.Error(err))
	}

	//start GRPC
	go serveGRPC(logger, db.NewStore(conn))

	// start HTTP
	serveHTTP(logger)
}

func serveGRPC(logger *zap.Logger, store *db.Store) {

	srv := server.NewRPCServer(logger, store)
	logger.Info("Starting GRPC Server...", zap.String("bind_addr", *bindAddr), zap.Int("bind_port", *grpcPort))
	logger.Fatal("GRPC Server Failed!", zap.Error(srv.Serve(*bindAddr, *grpcPort)))
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
	mux.Handle("/", http.FileServer(http.Dir(*webRoot)))
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", gwmux))
	mux.Handle("/metrics", promhttp.Handler())

	bind := fmt.Sprintf("%s:%d", *bindAddr, *httpPort)

	logger.Info("Starting HTTP...", zap.String("bind", bind))
	logger.Fatal("HTTP Failed", zap.Error(http.ListenAndServe(bind, mux)))
}
