package middleware

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Claims struct {
	Email string
	Role  string
	jwt.StandardClaims
}

type JwtClaims struct {
	cfg *config.Configure
}

func ValidateToken(ctx *gin.Context, cfg config.Configure, role string) (string, error) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		return "", errors.New("bearer token is missing, unable to proceed with request")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[7:])
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		return "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		return "", errors.New("token expired")
	}

	userRole := claims.Role
	if userRole != role {
		return "", errors.New("unauthorized user")
	}
	return claims.Email, nil
}

func ValidatePNRToken(ctx *gin.Context, cfg config.Configure, role string) (string, error) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		return "", errors.New("bearer token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[7:])
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		return "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		return "", errors.New("token expired")
	}
	userRole := claims.Role
	if userRole != role {
		return "", errors.New("unauthorized user")
	}
	return claims.Email, nil
}
