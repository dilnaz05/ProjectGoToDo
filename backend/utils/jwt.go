package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Token генерация
func GenerateJWT(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Token тексеру
func ParseJWT(tokenStr string) (uint, string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return 0, "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("could not parse claims")
	}

	userID := uint(claims["userID"].(float64))
	role := claims["role"].(string)

	// Debug шығару
	fmt.Println("Token string:", tokenStr)
	fmt.Println("Token valid:", token.Valid)
	fmt.Println("userID:", userID)
	fmt.Println("role:", role)

	return userID, role, nil
}
