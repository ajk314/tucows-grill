package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
    user := "root"
    password := "root"
	hostname := "localhost:3306"
    // hostname := "mysql-tucows-db:3306" // if not running locally, it would be this host
    dbname := "tucows"

    connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, hostname, dbname)
    db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// health check
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	} else {
		fmt.Println("Successfully connected to the database!")
	}
	return db, err
}
