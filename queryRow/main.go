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

	var name string
	var weight int64
	err = conn.QueryRow(ctx, "select name, weight from widgets where pk=$1", 2).Scan(&name, &weight)
	if err != nil {
		log.Fatal("QueryRow failed:", err)
	}

	fmt.Println(name, weight)
}
