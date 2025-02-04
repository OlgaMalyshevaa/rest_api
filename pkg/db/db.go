package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func Connect() *pgx.Conn {
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("Connection verification error: %v", err)
	}

	log.Println("Connected to database")
	return conn
}
