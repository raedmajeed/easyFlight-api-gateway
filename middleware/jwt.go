package middleware

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Claims struct {
	email string
	role  string
	jwt.StandardClaims
}

type JwtClaims struct {
	cfg *config.Configure
}

func ValidateToken(ctx *gin.Context, cfg config.Configure, role string) (string, error) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		log.Print("Header token missing")
		return "", errors.New("header token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[:7])
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return "", errors.New("error parsing token")
	}
	if !parserToken.Valid {
		log.Print("Invalid token")
		return "", errors.New("token invalid")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
		return "", errors.New("token expired")
	}

	userRole := claims.role
	if userRole != role {
		log.Println("Unauthorized user")
		return "", errors.New("unauthorized user")
	}
	return claims.email, nil
}
