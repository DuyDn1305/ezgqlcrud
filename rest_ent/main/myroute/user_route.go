package myroute

import (
	"context"
	"fmt"
	"restent/ent"
	"restent/ent/user"

	jwt "github.com/golang-jwt/jwt/v4"

	"github.com/labstack/echo/v4"
)

type userRoute _Route

var User = userRoute{}

// user
func (route *userRoute) getAll(c echo.Context) error {
	u, err := route.c.User.Query().All(route.ctx)
	if err != nil {
		return echo.NewHTTPError(400, "Bad request")
	}
	return c.JSON(200, u)
}

func (route *userRoute) getUser(c echo.Context) error {
	email := c.Param("email")
	claims := c.Get("claims").(Dict)
	if claims["email"] == email {
		u, err := route.c.User.Query().Where(user.Email(email)).Only(route.ctx)
		if err != nil {
			return badRequest(c, "User not found")
		}
		return c.JSON(200, Dict{"user": Dict{
			"name": u.Name,
			"email": u.Email,
		}})
	}
	return unAuthorized(c) 
}

func (route *userRoute) testMap(c echo.Context) error {
	abc := "asdfds"
	zzz := 3213
	vvv := []int{1, 2}
	return c.JSON(200, Dict{
		"abc": abc,
		"zzz": zzz,
		"vvv": vvv,
	})
}

func getClaims(tokenString string) Dict {
    new_token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return SecretKey, nil
    })
	
    if claims, ok := new_token.Claims.(jwt.MapClaims); new_token.Valid && ok {
		return claims;
	} 
	return nil
}


func (route *userRoute) Init(router *echo.Group, c *ent.Client, ctx context.Context) {
	route.c = c
	route.ctx = ctx
	router.Use(authMiddleware)
	router.GET("/all", route.getAll)
	router.GET("/get/:email", route.getUser)
	router.GET("/test", route.testMap)
}
