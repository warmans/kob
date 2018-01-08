package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/warmans/kob/server/pkg/search"

	"github.com/gocraft/dbr"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/warmans/kob/server/pkg/auth/oauth/google"
	"github.com/warmans/kob/server/pkg/auth/token"
	"github.com/warmans/kob/server/pkg/db"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/warmans/kob/server/pkg/rpc/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"github.com/warmans/kob/server/pkg/bot"
)

var (
	Version = "unknown"
)

var (
	bindAddr        = flag.String("bind-addr", "", "bind to this interface")
	httpPort        = flag.Int("http-port", 8080, "port to bind HTTP server")
	grpcPort        = flag.Int("grcp-port", 9090, "port to bind gRPC server")
	webRoot         = flag.String("web-root", "build/dist", "path to public dir")
	dbDSN           = flag.String("db-dsn", "", "DSN of postgres DB")
	host            = flag.String("host", "http://localhost:4200", "host of service inc. scheme. Used for oauth redirects")
	googleCreds     = flag.String("google-creds", "google-creds.json", "Google auth credentials")
	tokenKeyPath    = flag.String("token-key-path", "var/keys/jwtRS256.key", "Private RS256 RSA key used to create and validate tokens")
	searchIndexPath = flag.String("search-index-path", "var/entries.bleve", "Private RS256 RSA key used to create and validate tokens")
	slackToken = flag.String("slack-token", "", "Token used for slackbot integration")
	slackEnabled = flag.Bool("slack-enabled", false, "Enabled slack integration")
)

func main() {
	flag.Parse()

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// search index handles full text searching via bleve.
	searchIndex := makeSearchIndex(logger)

	// store handles permanent storeage of entities.
	store := makeStore(logger, searchIndex)

	// oauth handlers and verification strategy required by middlewares.
	oauthHandler, tokenStrategy := makeOauth(logger, store)

	// slackbot monitors slack for messages and suggests content
	if *slackEnabled {
		slackBot := makeSlackBot(logger, store, searchIndex)
		go func() {
			logger.Info("starting slack bot...")
			if err := slackBot.Run(); err != nil {
				logger.Error("slack bot failed", zap.Error(err))
				return
			}
		}()
	}

	// start servers
	go serveGRPC(logger, store, tokenStrategy, searchIndex)
	serveHTTP(logger, oauthHandler)
}

func serveGRPC(logger *zap.Logger, store *db.Store, tokenStrategy token.Strategy, searchIndex *search.Index) {

	srv := server.NewRPCServer(logger, store, tokenStrategy, searchIndex)
	logger.Info("Starting GRPC Server...", zap.String("bind_addr", *bindAddr), zap.Int("bind_port", *grpcPort))
	logger.Fatal("GRPC Server Failed!", zap.Error(srv.Serve(*bindAddr, *grpcPort)))
}

func serveHTTP(logger *zap.Logger, oauthHandler *google.OauthHandler) {

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
	mux.HandleFunc("/oauth/g/redirect", http.HandlerFunc(oauthHandler.HandleLoginRequest))
	mux.HandleFunc("/oauth/g/return", http.HandlerFunc(oauthHandler.HandleLoginResponse))

	bind := fmt.Sprintf("%s:%d", *bindAddr, *httpPort)

	logger.Info("Starting HTTP...", zap.String("bind", bind))
	logger.Fatal("HTTP Failed", zap.Error(http.ListenAndServe(bind, mux)))
}

func makeStore(logger *zap.Logger, index *search.Index) *db.Store {
	conn, err := dbr.Open("postgres", *dbDSN, nil)
	if err != nil {
		logger.Fatal("DB connection not possible", zap.Error(err))
	}
	store := db.NewStore(conn, index)

	//always re-index documents on startup. Or maybe don't. todo: figure out if it's a bad idea.
	go func() {
		err := store.WithSession(func(sess *db.Session) error {
			logger.Info("starting re-indexing...", zap.Error(err))
			startTime := time.Now()
			err := sess.Reindex()
			logger.Info("re-indexing complete...", zap.Float64("seconds_taken", time.Since(startTime).Seconds()))
			return err
		})
		if err != nil {
			logger.Error("Indexing failed", zap.Error(err))
		}
	}()

	return store
}

func makeOauth(logger *zap.Logger, store *db.Store) (*google.OauthHandler, token.Strategy) {
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
	oauthHandler := google.NewOauthHandler(
		oauthClient,
		cache.New(time.Hour, time.Minute),
		store,
		logger,
		tokenStrategy,
	)
	return oauthHandler, tokenStrategy
}

func makeSearchIndex(logger *zap.Logger) *search.Index {
	searchIndex, err := search.NewIndex(*searchIndexPath)
	if err != nil {
		logger.Fatal("Failed to create search index", zap.Error(err))
	}	
	return searchIndex
}

func makeSlackBot(logger *zap.Logger, store *db.Store, index *search.Index) *bot.SlackBot {
	return bot.NewSlackBot(*slackToken, logger, store, index)
}