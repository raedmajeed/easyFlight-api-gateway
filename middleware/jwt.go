package middleware

import (
	"errors"
	"fmt"
	"log"
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
		log.Print("bearer token missing")
		return "", errors.New("bearer token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[7:])
	fmt.Println(cfg.SECRETKEY)
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		log.Printf("error parsing token: %v", err)
		return "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		log.Print("invalid token")
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return "", errors.New("token expired")
	}

	userRole := claims.Role
	if userRole != role {
		log.Println("unauthorized user")
		return "", errors.New("unauthorized user")
	}
	return claims.Email, nil
}

func ValidatePNRToken(ctx *gin.Context, cfg config.Configure, role string) (string, error) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		log.Print("bearer token missing")
		return "", errors.New("bearer token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[7:])
	fmt.Println(cfg.SECRETKEY)
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		log.Printf("error parsing token: %v", err)
		return "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		log.Print("invalid token")
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return "", errors.New("token expired")
	}
	userRole := claims.Role
	if userRole != role {
		log.Println("unauthorized user")
		return "", errors.New("unauthorized user")
	}
	return claims.Email, nil
}
