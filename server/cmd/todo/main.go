
package main

import (
    "context"
    "log"
    "net/http"

    "todo"
    "todo/ent"
    "todo/ent/migrate"

    // "entgo.io/ent/dialect/sql"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"

    // _ "github.com/mattn/go-sqlite3"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
    // Create ent.Client and run the schema migration.
    // client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	client, err := ent.Open("postgres","host=localhost port=5432 user=postgres dbname=testentgql password=root sslmode=disable")
    // client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/testentgql?parseTime=True")
    
    if err != nil {
        log.Fatal("opening ent client", err)
    }
    if err := client.Schema.Create(
        context.Background(),
        migrate.WithGlobalUniqueID(true),
    ); err != nil {
        log.Fatal("opening ent client", err)
    }


    // Configure the server and start listening on :8081.
    srv := handler.NewDefaultServer(todo.NewSchema(client))
	cors_handler := cors.AllowAll().Handler(srv)
    http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", cors_handler)
	log.Println("listening on :8081")
	

    if err := http.ListenAndServe(":8081", nil); err != nil {
        log.Fatal("http server terminated", err)
    }
}