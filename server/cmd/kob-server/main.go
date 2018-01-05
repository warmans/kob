package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/gocraft/dbr"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/warmans/kob/server/pkg/auth/oauth/google"
	"github.com/warmans/kob/server/pkg/auth/token"
	"github.com/warmans/kob/server/pkg/db"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/warmans/kob/server/pkg/rpc/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/patrickmn/go-cache"
)

var (
	Version = "unknown"
)

var (
	bindAddr     = flag.String("bind-addr", "", "bind to this interface")
	httpPort     = flag.Int("http-port", 8080, "port to bind HTTP server")
	grpcPort     = flag.Int("grcp-port", 9090, "port to bind gRPC server")
	webRoot      = flag.String("web-root", "build/dist", "path to public dir")
	dbDSN        = flag.String("db-dsn", "", "DSN of postgres DB")
	host         = flag.String("host", "http://localhost:4200", "host of service inc. scheme. Used for oauth redirects")
	googleCreds  = flag.String("google-creds", "google-creds.json", "Google auth credentials")
	tokenKeyPath = flag.String("token-key-path", "keys/jwtRS256.key", "Private RS256 RSA key used to create and validate tokens")
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	conn, err := dbr.Open("postgres", *dbDSN, nil)
	if err != nil {
		logger.Fatal("DB connection not possible", zap.Error(err))
	}

	store := db.NewStore(conn)

	key, err := token.ReadKeyFile(*tokenKeyPath)
	if err != nil {
		logger.Fatal("Failed to load token key", zap.Error(err))
	}
	tokenStrategy := token.NewJWTStrategy(key)

	// google credentials via https://console.developers.google.com/apis/credentials
	oauthClient, err := google.NewClient(*googleCreds, *host+"/oauth/g/return")
	if err != nil {
		logger.Fatal("Failed to create OAUTH client", zap.Error(err))
	}
	oauthHandlers := google.NewOauthHandler(
		oauthClient,
		cache.New(time.Hour, time.Minute),
		store,
		logger,
		tokenStrategy,
	)

	//start GRPC
	go serveGRPC(logger, store, tokenStrategy)

	// start HTTP
	serveHTTP(logger, oauthHandlers)
}

func serveGRPC(logger *zap.Logger, store *db.Store, tokenStrategy token.Strategy) {

	srv := server.NewRPCServer(logger, store, tokenStrategy)
	logger.Info("Starting GRPC Server...", zap.String("bind_addr", *bindAddr), zap.Int("bind_port", *grpcPort))
	logger.Fatal("GRPC Server Failed!", zap.Error(srv.Serve(*bindAddr, *grpcPort)))
}

func serveHTTP(logger *zap.Logger, oauthHandlers *google.OauthHandler) {

	grpcConn, err := grpc.Dial(fmt.Sprintf("127.0.0.1:%d", *grpcPort), grpc.WithInsecure())
	if err != nil {
		logger.Fatal("Failed to dial GRPC Server", zap.Error(err))
	}
	defer grpcConn.Close()

	gwmux := runtime.NewServeMux()
	rpc.RegisterKobServiceHandler(context.Background(), gwmux, grpcConn)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*webRoot)))
	mux.Handle("/api/", gwmux)
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/oauth/g/redirect", http.HandlerFunc(oauthHandlers.HandleLoginRequest))
	mux.HandleFunc("/oauth/g/return", http.HandlerFunc(oauthHandlers.HandleLoginResponse))

	bind := fmt.Sprintf("%s:%d", *bindAddr, *httpPort)

	logger.Info("Starting HTTP...", zap.String("bind", bind))
	logger.Fatal("HTTP Failed", zap.Error(http.ListenAndServe(bind, mux)))
}
