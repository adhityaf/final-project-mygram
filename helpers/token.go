package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "secret"

func GenerateToken(id, age uint, email, username string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"age": age,
		"email": email,
		"username": username,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}

func VerifyToken(tokenStr string) (interface{}, error) {
	errResponse := errors.New("token invalid")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
