package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// Change this connection string according to your MySQL setup
	dsn := "root:Impelsys@2020@tcp(127.0.0.1:3306)/inventorydb"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database: ", err)
	}

	fmt.Println("Connected to the database successfully!")
}
