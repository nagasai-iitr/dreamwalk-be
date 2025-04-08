package main

import (
	"context"
	//"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func InitDB() {
	var err error
	dbURL := os.Getenv("DB_URL")
	conn, err = pgx.Connect(context.Background(), dbURL)

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS contracts (
			id TEXT PRIMARY KEY,
			min_amount FLOAT,
			max_amount FLOAT,
			receiver_address TEXT,
			status TEXT
		)
	`)
	if err != nil {
		log.Fatal("Failed to initialize schema:", err)
	}
}

func InsertEntry(c Entry) error {
	_, err := conn.Exec(context.Background(),
		`INSERT INTO contracts (id, min_amount, max_amount, receiver_address, status) VALUES ($1, $2, $3, $4, $5)`,
		c.ID, c.MinAmount, c.MaxAmount, c.ReceiverAddress, c.Status)
	return err
}

func GetInitEntries() ([]Entry, error) {
	rows, err := conn.Query(context.Background(), `SELECT id, min_amount, max_amount, receiver_address, status FROM contracts WHERE status = 'INIT'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []Entry
	for rows.Next() {
		var c Entry
		if err := rows.Scan(&c.ID, &c.MinAmount, &c.MaxAmount, &c.ReceiverAddress, &c.Status); err != nil {
			return nil, err
		}
		entries = append(entries, c)
	}
	return entries, nil
}

func UpdateEntryStatus(id, status string) {
	_, err := conn.Exec(context.Background(), `UPDATE contracts SET status=$1 WHERE id=$2`, status, id)
	if err != nil {
		log.Printf("Error updating contract %s: %v", id, err)
	}
}
