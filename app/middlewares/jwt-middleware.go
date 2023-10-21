package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func JWTMiddleware() echo.MiddlewareFunc {
	val, exist := os.LookupEnv("SECRETJWT")

	if !exist {
		logrus.Error("missing secret jwt")
	}
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(val),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId string, role string) (string, error) {
	val, exist := os.LookupEnv("SECRETJWT")

	if !exist {
		logrus.Error("missing secret jwt")
	}
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(val))

}

func ExtractTokenUserId(e echo.Context) int {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(float64)
		return int(userId)
	}
	return 0
}
