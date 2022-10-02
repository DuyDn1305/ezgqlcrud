package myroute

import (
	"context"
	"restent/ent"
	"strings"

	"github.com/labstack/echo/v4"
)

type _Route struct {
	c   *ent.Client
	ctx context.Context
}

func unAuthorized(c echo.Context) error {
	return c.JSON(401, Dict{"message": "Unauthorized"})
}

func badRequest(c echo.Context, message string) error {
	return c.JSON(400, Dict{"message": message})
}

func severError(c echo.Context, message string) error {
	return c.JSON(500, Dict{"message": message})
}

type Dict = map[string]interface{}

var SecretKey = []byte("very very secret") // lowercase

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func (c echo.Context) error {
		// if valid
		if tokenStrings, ok := c.Request().Header["Authentication"]; ok {
			token := strings.Split(tokenStrings[0], " ")[1]
			if claims := getClaims(token); claims != nil {
				// fmt.Println(token)
				c.Set("claims", claims)
				return next(c)
			}
		}
		return c.JSON(401, Dict{
			"message": "Unauthorized",
		})
	}
}  