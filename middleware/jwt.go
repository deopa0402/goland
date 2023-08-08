package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"main/domain"
	"main/service"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authHeader := c.Request().Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}

		tokenString := parts[1]

		token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return service.SecretKey, nil
			})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
		if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
			c.Set("id", claims.ID)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}
		if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
			c.Set("id", claims.ID)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		return next(c)
	}
}
