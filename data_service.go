package main

import (
	"database/sql"
)

var conn *sql.DB

func init() {
	conn = connect()
	createDB(conn)

}

func AddData(service service) {
	insertData(conn, service)
}
