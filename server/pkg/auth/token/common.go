package token 

import (
	"context"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"go.uber.org/zap"
)

type Strategy interface{
	CreateToken(author *rpc.Author) (string, error)
	ValidateToken(tokenString string) (*rpc.Author, error)
} 

type authorCtx string

const authorKey = authorCtx("author")

func ExtractAuthor(ctx context.Context) (*rpc.Author) {
	return ctx.Value(authorKey).(*rpc.Author)
}

func AuthFunc(tokenStrategy Strategy, logger *zap.Logger) (grpc_auth.AuthFunc) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			logger.Debug("failed to find token in request", zap.Error(err))
			return nil, status.Error(codes.PermissionDenied, "token not found in request")
		}
		author, err := tokenStrategy.ValidateToken(token)
		if err != nil {
			logger.Debug("invalid token encountered", zap.Error(err))
			return nil, status.Error(codes.PermissionDenied, "token was invalid")
		}
		return context.WithValue(ctx, authorKey, author), nil
	}
}
