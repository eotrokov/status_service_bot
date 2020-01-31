package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go"
)

func createDB(connect *sql.DB) {
	_, err := connect.Exec(`
		CREATE TABLE IF NOT EXISTS services (
			url String,
			status       UInt8,
			action_time  DateTime
		) engine=Memory
	`)

	if err != nil {
		log.Fatal(err)
	}
}

func insertData(connect *sql.DB, service service) {
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO services (url, status, action_time) VALUES (?, ?, ?)")
	)
	defer stmt.Close()

	if _, err := stmt.Exec(
		service.url,
		service.status,
		time.Now(),
	); err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func connect() *sql.DB {
	connStr := os.Getenv("clickhouse")
	connect, err := sql.Open("clickhouse", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
		return nil
	}

	return connect
}
