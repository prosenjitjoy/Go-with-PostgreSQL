package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(ctx)

	tx, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal("Unable to begin transaction:", err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "insert into foo(id) values (1)")
	if err != nil {
		log.Fatal("Unable to execute query:", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatal("Unable to commit query:", err)
	}

	err = pgx.BeginFunc(ctx, conn, func(tx pgx.Tx) error {
		_, err := tx.Exec(ctx, "insert into foo(id) values (1)")
		return err
	})
	if err != nil {
		log.Fatal("Transaction failed")
	}
}
