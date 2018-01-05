package token 

import (
	"context"
	"github.com/warmans/kob/server/pkg/rpc"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
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

func AuthFunc(tokenStrategy Strategy) (grpc_auth.AuthFunc){
	return func(ctx context.Context) (context.Context, error) {
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}
		author, err := tokenStrategy.ValidateToken(token)
		if err != nil {
			return nil, err
		}
		return context.WithValue(ctx, authorKey, author), nil
	}
}
