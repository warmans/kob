package token

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/warmans/kob/server/pkg/rpc"
)

type Claims struct {
	jwt.StandardClaims
	Author *rpc.Author `json:"author"`
}

func NewJWTStrategy(k *rsa.PrivateKey) *JWTStrategy {
	return &JWTStrategy{privateKey: k}
}

type JWTStrategy struct {
	privateKey *rsa.PrivateKey
}

func (t *JWTStrategy) CreateToken(author *rpc.Author) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &Claims{Author: author})
	return token.SignedString(t.privateKey)
}

func (t *JWTStrategy) ValidateToken(tokenString string) (*rpc.Author, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return &t.privateKey.PublicKey, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims.Author, nil
}

func ReadKeyFile(path string) (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	privateKey, err := ioutil.ReadAll(privateKeyFile)
	if err != nil {
		return nil, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(privateKey)
}
