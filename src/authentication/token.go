package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken create authentication token
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken check that the token in the request is valid
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, getVerificationKey)
	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

// ExtractUserID extract the user id from the token
func ExtractUserID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, getVerificationKey)
	if error != nil {
		return 0, error
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, error := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)
		if error != nil {
			return 0, nil
		}

		return userID, nil
	}

	return 0, errors.New("Invalid token")
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Sign method is unexpected! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	splitToken := strings.Split(token, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
