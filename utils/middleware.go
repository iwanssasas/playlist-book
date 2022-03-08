package utils

import (
	"PLAYLISTBOOK/config"
	"PLAYLISTBOOK/errs"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ExtractClaims(secret, tokenStr string) (jwt.MapClaims, error) {
	hmacSecret := []byte(secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid JWT Token")
}

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.Get()
		token := c.GetHeader("Authorization")
		parts := strings.Split(token, " ")
		if token == "" {
			err := errs.ErrUnauthorized
			c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			c.Abort()
			return
		}
		claims, err := ExtractClaims(conf.Secret, parts[1])
		if err != nil {
			err := errs.ErrUnauthorized
			c.JSON(http.StatusUnauthorized, ErrorResponse(err))
			c.Abort()
			return
		}
		data := claims["Data"]
		result := map[string]interface{}{}
		encoded, _ := json.Marshal(data)
		json.Unmarshal(encoded, &result)
		for key, val := range result {
			c.Set(key, val)
		}
		c.Next()
	}
}
