package main

import (
	"context"
	"log"
	"restent/ent"
	_ "restent/ent/runtime"
	"restent/main/myroute"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

// func main() {
//     // sample token string taken from the New example

//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
//         "foo": "hfghghghg",
//         "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
//     })

//     // Sign and get the complete encoded token as a string using the secret
//     tokenString, err := token.SignedString(hmacSampleSecret)

//     fmt.Println(tokenString, err)

//     // old_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.-BRTwjN-sAlUjO-82qDrNHdMtGAwgWH05PrN49Ep_sU"

//     // tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"
//     // // Parse takes the token string and a function for looking up the key. The latter is especially
//     // // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
//     // // head of the token to identify which key to use, but the parsed token (head and claims) is provided
//     // // to the callback, providing flexibility.
//     new_token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         // Don't forget to validate the alg is what you expect:
//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//         }

//         // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
//         return hmacSampleSecret, nil
//     })

//     if claims, ok := new_token.Claims.(jwt.MapClaims); ok && new_token.Valid {
//         fmt.Println(claims["foo"], claims["nbf"])
//     } else {
//         fmt.Println(ok, new_token.Valid)
//     }

// }

func main() {
    client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=restent password=root sslmode=disable")
    if err != nil {
        log.Fatalf("failed opening connection to postgres: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    ctx := context.Background()

    e := echo.New()
    // e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"},
    }))

    myroute.User.Init(e.Group("/users"), client, ctx)
    myroute.Blog.Init(e.Group("/blogs"), client, ctx)
    myroute.NoAuth.Init(e.Group("/noauth"), client, ctx)
    e.Logger.Fatal(e.Start(":8080"))
}