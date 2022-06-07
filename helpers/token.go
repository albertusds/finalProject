package helpers

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "secret"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := parseToken.SignedString([]byte(SECRET_KEY))

	return signedToken
}

func VerifyToken(tokenStr string) (interface{}, error) {

	errResp := errors.New("token invalid")
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, nok := t.Method.(*jwt.SigningMethodHMAC); !nok {
			return nil, errResp
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if _, nok := token.Claims.(jwt.MapClaims); !nok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}
