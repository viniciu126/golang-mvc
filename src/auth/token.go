package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Create a JWT with user permissions
func CreateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

// ValidateToken validate if request token is valid
func ValidateToken(r *http.Request) error {
	tokenString := getToken(r)

	token, err := jwt.Parse(tokenString, getTokenVerificationKey)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token invalid")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	tokenSplit := strings.Split(token, " ")

	if len(tokenSplit) == 2 {
		return tokenSplit[1]
	}

	return ""
}

func getTokenVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Sign method unexpected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

// GetUserID returns userID stored in request token
func GetUserID(r *http.Request) (uint64, error) {
	tokenString := getToken(r)

	token, err := jwt.Parse(tokenString, getTokenVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return userID, nil
	}

	return 0, errors.New("Token inválido")
}
