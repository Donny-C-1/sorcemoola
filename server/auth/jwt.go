
package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/donny-c-1/sorcemoola/server/models"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

func InitJWT() error {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return fmt.Errorf("JWT_SECRET environement variable is not set")
	}
	JWTSecret = []byte(secret)
	return nil
}

func GenerateToken(user *models.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWTSecret)

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
	if len(JWTSecret) == 0 {
		log.Println("JWT Secret not set")
		return nil, errors.New("JWT secret not set")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ExtractUserID(tokenStr string) (string, error) {
	claims, err := ValidateToken(tokenStr)
	if err != nil {
		return "", err
	}

	if userID, ok := claims["user_id"].(string); ok {
		return userID, nil
	}

	return "", errors.New("user_id not found in token")
}
