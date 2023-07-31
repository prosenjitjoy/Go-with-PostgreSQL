package main

import (
	"context"
	"fmt"
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

	var sum, n int
	rows, err := conn.Query(ctx, "select generate_series(1, $1)", 10)
	if err != nil {
		log.Fatal("Unable to execute query:", err)
	}
	tag, err := pgx.ForEachRow(rows, []any{&n}, func() error {
		sum += n
		return nil
	})
	if err != nil {
		log.Fatal("Unable to execute ForEachRow:", err)
	}

	fmt.Println(tag)
	fmt.Println(n)
	fmt.Println(sum)
}
