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

	commandTag, err := conn.Exec(ctx, "delete from widgets where pk=$1", 2)
	if err != nil {
		log.Fatal("Unable to execute query:", err)
	}
	if commandTag.RowsAffected() != 1 {
		log.Fatal("No row found to delete!")
	} else {
		fmt.Println("Deleted successfully")
	}
}
