package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/raedmajeed/api-gateway/pkg/config"
)

type Claims struct {
	email		string
	role		string
	jwt.StandardClaims
}

func ValidateToken(ctx *gin.Context, cfg config.Configure, role string) {
	headerToken := ctx.GetHeader("Authorization")
	if headerToken == "" {
		log.Print("Header token missing")
	}

	claims := &Claims{}
	token := string([]byte(headerToken)[:7])
	parserToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SECRETKEY), nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
	}
	if !parserToken.Valid {
		log.Print("Invalid token")
	}

	expTime := claims.ExpiresAt
	if expTime < time.Now().Unix() {
		log.Print("token Expired")
	}

	userRole := claims.role
	if userRole != role {
		log.Println("Unauthorized user")
	}
	ctx.Set("userEmail", claims.email)
	ctx.Next()
}