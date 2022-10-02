package main

import (
	"context"
	"log"
	"restent/ent"
	"restent/main/pkg"

	_ "github.com/lib/pq"
)

// func dropDB() {
// 	pkg.PrintC(pkg.ColorGreen, "Connecting to postgress...")
// 	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=root sslmode=disable")
// 	if err != nil {
// 		log.Fatalf(pkg.Colorize(pkg.ColorRed, "Failed opening connection to postgres: ")+ "%v", err)
// 	}
// 	defer client.Close()

// }

func Build() {
	pkg.PrintC(pkg.ColorGreen, "Connecting to postgress...")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=restent password=root sslmode=disable")
	if err != nil {
		log.Fatalf(pkg.Colorize(pkg.ColorRed, "Failed opening connection to postgres: ")+ "%v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf(pkg.Colorize(pkg.ColorRed, "failed creating schema resources: ")+"%v", err)
	}
	pkg.PrintC(pkg.ColorGreen, "Schema created!")

}

func main() {
	Build();
}