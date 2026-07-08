package middleware

import (
	"net/http"
	"strings"

	"tetra-server/config"

	"github.com/gin-gonic/gin"

	"errors"
	"tetra-server/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID   uint        `json:"user_id"`
	Username string      `json:"username"`
	Role     models.Role `json:"role"`

	jwt.RegisteredClaims
}

func GenerateAccessToken(user *models.User, cfg *config.Config) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.AppName,
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(cfg.JWTAccessExpire) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWTSecret))
}

func GenerateRefreshToken(user *models.User, cfg *config.Config) (string, error) {

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   cfg.AppName,
			Subject:  user.Username,
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					time.Duration(cfg.JWTRefreshExpire) * time.Hour,
				),
			),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.JWTSecret))
}

func ParseToken(tokenString string, cfg *config.Config) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}

			return []byte(cfg.JWTSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header is required",
			})
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(header, "Bearer ")

		claims, err := ParseToken(tokenString, cfg)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
