package main

import (
	"context"
	"fmt"
	"log"

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

	var name string
	var weight int64
	err = dbpool.QueryRow(ctx, "select name, weight from widgets where pk=$1", 1).Scan(&name, &weight)
	if err != nil {
		log.Fatal("QueryRow failed:", err)
	}

	fmt.Println(name, weight)
}
