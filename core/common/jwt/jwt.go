package jwt

import (
	"fmt"
	"github.com/wsw365904/wswlog/wlogging"

	"github.com/dgrijalva/jwt-go"
)

var logger = wlogging.MustGetLoggerWithoutName()

func CreateToken(key string, m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
