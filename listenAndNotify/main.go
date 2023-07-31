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

	_, err = conn.Exec(ctx, "listen channelname")
	if err != nil {
		log.Fatal("Unable to listen from channelname")
	}

	notification, err := conn.WaitForNotification(ctx)
	if err != nil {
		log.Fatal("Unable to listen for notification")
	}

	// do something with the notification
	fmt.Println(notification)
}
