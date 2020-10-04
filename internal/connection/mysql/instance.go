package mysql

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"

	// Register mysql
	_ "github.com/go-sql-driver/mysql"
)

func Open() *sql.DB {
	var conn *sql.DB
	var err error

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?collation=utf8mb4_unicode_ci",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	if conn, err = sql.Open("mysql", dsn); err != nil {
		log.Fatal(err)
	}

	return conn
}

func Close(db io.Closer) {
	_ = db.Close()
}
