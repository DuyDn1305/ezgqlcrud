package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"restent/ent"
	_ "restent/ent/runtime"
	"restent/main/pkg"

	_ "github.com/lib/pq"
)

func rollback(tx *ent.Tx, err error) error {
    if rerr := tx.Rollback(); rerr != nil {
        err = fmt.Errorf("%w: %v", err, rerr)
    }
    return err
}

func withTrans(client *ent.Client, ctx context.Context, fn func (*ent.Client, context.Context) error) error {
    tx, err := client.Tx(ctx)
    if err != nil {
		fmt.Println("cannot create tx: ", err)
		return err
    }
    txClient := tx.Client()
	if err := fn(txClient, ctx); err != nil {
        rollback(tx, err)
		return err
    }
	return tx.Commit()
}

func populate(client *ent.Client, ctx context.Context) error {
	// cate
	client.Cate.Create().SetName("WFH").SaveX(ctx)
	nat := client.Cate.Create().SetName("Natural").SaveX(ctx)
	client.Cate.Create().SetName("Town").SaveX(ctx)
	human := client.Cate.Create().SetName("Human").SaveX(ctx)
	client.Cate.Create().SetName("Technology").SaveX(ctx)

	// user
	duy := client.User.Create().
		SetEmail("duydn1305@gmail.com").SetName("Duy Đn").SetPassword("123456").SaveX(ctx)
	vu := client.User.Create().
	SetEmail("riz@gmail.com").SetName("Vu NA").SetPassword("123456").SaveX(ctx)

	// blog
	bl := client.Blog.Create().
		SetAuthor(duy).
		SetTitle("Best photo for natural").
		SetContent("<div>Header1</div><div>Content1</div>").
		SetThumbnail("https://phlearn.com/wp-content/uploads/2020/01/annie-spratt-ogDort6vKuE-unsplash.jpg").
		AddTags(nat, human).
		SaveX(ctx)
	
	// comment
	c := client.Comment.Create().
		SetContent("Where is this place?").
		SetWriter(vu).
		SetBelongto(bl).
		SaveX(ctx)

	
	client.Comment.Create().
		SetContent("It is from Cabin!").
		SetWriter(duy).
		SetReplyTo(c).
		SetBelongto(bl).
		SaveX(ctx)

	return nil
}

func delete_all(client *ent.Client, ctx context.Context) error {
	client.Comment.Delete().ExecX(ctx)
	client.Blog.Delete().ExecX(ctx)
	client.User.Delete().ExecX(ctx)
	client.Cate.Delete().ExecX(ctx)
	return nil
}

func main() {
	pkg.PrintC(pkg.ColorBlue, "Connecting to postgress...")
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=restent password=root sslmode=disable")
	if err != nil {
		log.Fatalln("Cannot connect: "+pkg.Colorize(pkg.ColorRed, err.Error()))
	}

	// Run the auto migration tool.
	defer client.Close()
	ctx := context.Background()
	
	replace := flag.Bool("replace", false, "Restart database and init with seeding")
	flag.Parse()

	if *replace {
		pkg.PrintC(pkg.ColorYellow, "Deleting old data....")
		if  err := withTrans(client, ctx, delete_all); err != nil {
			log.Fatalln("Delete failed: "+pkg.Colorize(pkg.ColorRed, err.Error()))
		}
	}

	pkg.PrintC(pkg.ColorBlue, "Populating sample...")
	if  err := withTrans(client, ctx, populate); err != nil {
		log.Fatalln("Seeding failed: "+pkg.Colorize(pkg.ColorRed, err.Error()))
	}

}

