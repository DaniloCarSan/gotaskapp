package security

import (
	"errors"
	"fmt"
	"gotaskapp/src/config"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// Generate token jwt
func GenerateJwtToken(id uint64, duration time.Duration) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(duration).Unix()
	permissions["id"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.JWT_SECRET))
}

// Validate jwt token
func ValidateJwtToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, funcSecretJwtKeyVerifySignature)

	if err != nil {
		return &jwt.Token{}, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	}

	return &jwt.Token{}, errors.New("token invalid")
}

func ExtractJwtTokenFromHeaderAuthorization(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) != 2 {
		return "", errors.New("bearer or token not found in authorization")
	}

	return strings.Split(token, " ")[1], nil
}

func funcSecretJwtKeyVerifySignature(token *jwt.Token) (interface{}, error) {

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.JWT_SECRET), nil
}

func ExtractUserIdOfJwtToken(token *jwt.Token) (uint64, error) {

	permissions, _ := token.Claims.(jwt.MapClaims)

	id, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)

	if err != nil {
		return 0, err
	}

	return id, nil
}
