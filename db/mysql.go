package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() {
	err := db.Close()
	if err != nil {
		log.Println("Error closing the database connection:", err)
	} else {
		log.Println("Closed the database connection")
	}
}
