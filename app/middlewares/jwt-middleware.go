package middlewares

import (
	"mini_project/app/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type TokenData struct {
	Id   string
	Role string
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(configs.SECRET_JWT),
		SigningMethod: "HS256",
	})
}

func CreateToken(userId string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.SECRET_JWT))

}

func ExtractToken(e echo.Context) TokenData {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := claims["userId"].(string)
		role := claims["role"].(string)
		return TokenData{
			Id:   userId,
			Role: role,
		}
	}
	return TokenData{
		Id:   "",
		Role: "",
	}
}
