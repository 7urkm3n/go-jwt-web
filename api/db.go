package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func connectDb() (*sql.DB, error) {
	// delete the file to avoid duplicated records.
	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	if err != nil {
		return nil, err
	}

	createTable(sqliteDatabase)
	insertUser(sqliteDatabase, "rr@rr.com", "yyyy")
	insertUser(sqliteDatabase, "rr1@rr.com", "xxxx")

	return sqliteDatabase, nil
}

// createTable(sqliteDatabase) // Create Database Tables
func createTable(db *sql.DB) {
	createUserTableSQL := `CREATE TABLE users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"email" TEXT NOT NULL UNIQUE,
		"password" TEXT,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);` // SQL Statement for Create Table

	log.Println("Create user table...")
	statement, err := db.Prepare(createUserTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertUser(db *sql.DB, email, password string) {
	log.Println("Inserting user record ...")
	insertUserSQL := `INSERT INTO users(email, password) VALUES (?, ?)`
	statement, err := db.Prepare(insertUserSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(email, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getUsers(db *sql.DB) {
	row, err := db.Query("SELECT * FROM users ORDER BY email")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var email string
		var password string
		row.Scan(&id, &email, &password)
		log.Println("User: ", email, " ", password)
	}
}
