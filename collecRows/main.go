package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Unable to create connection pool:", err)
	}
	defer dbpool.Close()

	rows, err := dbpool.Query(ctx, "select generate_series(1, $1)", 5)
	if err != nil {
		log.Fatal("Unable to execute query:", err)
	}
	numbers, err := pgx.CollectRows(rows, pgx.RowTo[int])
	if err != nil {
		log.Fatal("Unable to collect rows:", err)
	}

	fmt.Println(numbers)
}
