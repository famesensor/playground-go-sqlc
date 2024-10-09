package main

import (
	"context"
	"log"

	"github.com/famesensor/playground-go-sqlc/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	d, err := pgxpool.New(ctx, "user=username password=pass dbname=local_db sslmode=disable host=localhost port=5432 search_path=sqlc")
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(d)

	insertedAuthor, err := q.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Mr Beans",
		Bio: pgtype.Text{
			String: "The funniest man alive",
			Valid:  true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(insertedAuthor)

	tx, err := d.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	insertedAuthor2, err := q.WithTx(tx).CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Mr Beans 2",
		Bio: pgtype.Text{
			String: "The funniest man alive",
			Valid:  true,
		},
	})
	if err != nil {
		tx.Rollback(ctx)
		log.Fatal(err)
	}

	insertedAuthor3, err := q.WithTx(tx).CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Mr Beans 2",
		Bio: pgtype.Text{
			String: "The funniest man alive",
			Valid:  true,
		},
	})
	if err != nil {
		tx.Rollback(ctx)
		log.Fatal(err)
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println(insertedAuthor2)
	log.Println(insertedAuthor3)
}
