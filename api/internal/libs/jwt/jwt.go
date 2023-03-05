package jwt

import (
	"api/internal/config"
	"api/internal/libs/constant"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var conf = config.GetConfig()

// Parses the jwt token
func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(constant.INVALID_SIGNING_METHOD)
		}

		return []byte(conf.JWT_AT_SECRET), nil
	})

	if err != nil {
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return jwt.MapClaims{}, nil
}
