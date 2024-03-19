package db

import (
	"github.com/go-pg/pg/v9"
	"log"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     getEnv("POSTGRES_USER", "admin"),
		Password: getEnv("POSTGRES_PASSWORD", "password"),
		Addr:     getEnv("POSTGRES_HOST", "localhost") + ":" + getEnv("POSTGRES_PORT", "5432"),
		Database: getEnv("POSTGRES_DB", "db"),
		OnConnect: func(conn *pg.Conn) error {
			return nil
		},
	}

	log.Println(opts)

	var db = pg.Connect(opts)

	if db == nil {
		log.Printf("Failed to connect to PotsgreSQL")
		os.Exit(100)
	}

	log.Printf("Connected to PostgreSQL")

	return db
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
