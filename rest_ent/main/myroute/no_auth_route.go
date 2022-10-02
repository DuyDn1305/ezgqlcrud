package myroute

import (
	"context"
	"restent/ent"
	"restent/ent/user"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type noauthRoute _Route

var NoAuth = noauthRoute{}

func create_token(email string, id uuid.UUID) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "exp": time.Now().Add(3*24*time.Hour).Unix(),
		// "iat": time.Now(),
        "email": email,
		"user_id": id.String(),

    })
    // Sign and get the complete encoded token as a string using the secret
    tokenString, _ := token.SignedString(SecretKey)
	return tokenString
}

func (route *noauthRoute) login(c echo.Context) error {
	var userInput ent.User

	if c.Bind(&userInput) == nil {
		u, err := route.c.User.Query().Where(user.Email(userInput.Email)).Only(route.ctx)
		if err == nil {

			if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userInput.Password)) == nil {
				// create new token, append to header
				return c.JSON(200, Dict{"token": create_token(userInput.Email, u.ID)})
			}
		}
	}
	return badRequest(c, "Login failed")

}


func (route *noauthRoute) register(c echo.Context) error {
	var userInput ent.User
	
	if c.Bind(&userInput) == nil {
		client := route.c
		_, err := client.User.Create().
		SetEmail(userInput.Email).
		SetName(userInput.Name).
		SetPassword(userInput.Password).Save(route.ctx)
		if err == nil {
			return c.JSON(200, Dict{
				"message": "success",
			})
		}
	}
	return badRequest(c, "Email existed")
}


func (route *noauthRoute) Init(router *echo.Group, c *ent.Client, ctx context.Context) {
	route.c = c
	route.ctx = ctx
	router.POST("/login", route.login)
	router.POST("/register", route.register)
}
