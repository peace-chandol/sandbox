package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/peace/sandbox/internal/models"
)

func GenerateToken(id string, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	secret := os.Getenv("JWT_SECRET_KEY")
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error in generating JWT", err)
		return "", err
	}
	return tokenStr, nil
}

func ParseToken(tokenStr string) (*models.User, error) {
	secret := os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	idStr, ok := claims["id"].(string)
	if !ok {
		return nil, err
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:    id,
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}

	return user, nil
}
