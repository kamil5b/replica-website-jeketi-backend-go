package util

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//THIS PACKAGE HAVE NO LOCAL IMPORT

func HashAndSalt(pwd []byte) string {
	hash, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash)
}
func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	return err == nil
}

func GenerateJWT(userID uuid.UUID, email, role string) (string, error) {
	tokenExpiryHour := GetEnv("TOKEN_EXPIRY_HOUR")
	expiryHour, err := strconv.Atoi(tokenExpiryHour)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"role":  role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(expiryHour)).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	secretKey := GetEnv("TOKEN_SECRET_KEY")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
