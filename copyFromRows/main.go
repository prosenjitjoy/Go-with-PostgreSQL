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

	rows := [][]any{
		{"John", 36},
		{"Jane", 29},
	}

	copyCount, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{"widgets"},
		[]string{"name", "weight"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		log.Fatal("Unable to copyFromRows:", err)
	}

	fmt.Println(copyCount)

}
