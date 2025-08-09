package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	jwt.StandardClaims
}

func GenerateAllTokens(id string, firstName string, lastName string, email string) (string, string, error) {
	key := os.Getenv("JWT_SECRET")
	claims := &SignedDetails{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))

	if err != nil {
		return " ", " ", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(key))

	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func ValidateToken(clientToken string) (SignedDetails, string) {
	secretKey := os.Getenv("JWT_SECRET")
	token, err := jwt.ParseWithClaims(clientToken, &SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
	var msg string

	if err != nil {
		msg = err.Error()
		return SignedDetails{}, msg
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "token is invalid"
		return SignedDetails{}, msg
	}

	if claims.ExpiresAt < time.Now().Unix() {
		msg = "token is expired"
		return SignedDetails{}, msg
	}

	return *claims, msg
}
