package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Widgets struct {
	Name   string
	Weight int
}

const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(ctx)

	rows := []Widgets{
		{Name: "Smith", Weight: 62},
		{Name: "Doe", Weight: 92},
	}

	copyCount, err := conn.CopyFrom(
		ctx,
		pgx.Identifier{"widgets"},
		[]string{"name", "weight"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			return []any{rows[i].Name, rows[i].Weight}, nil
		}),
	)

	fmt.Println(copyCount)
}
