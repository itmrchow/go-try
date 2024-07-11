package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	CREATE_USER_SCHEMA = `
		CREATE TABLE Users (
			user_id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
		`

	CREATE_POST_SCHEMA = `
		CREATE TABLE Posts (
			post_id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES Users(user_id)
		)
		`
)

func main() {

	db, err := connectDb()
	if err != nil {
		log.Fatalln("DB error: " + err.Error())
	}
	log.Println("connected db")

	if err := createSchema(db); err != nil {
		log.Fatalln("DB error: " + err.Error())
	}

	// create table

	checkTableExists(db, "users")

}

func createSchema(db *sqlx.DB) error {
	// check sqlx_test_db exists
	schemaExist, checkExistsErr := checkSchemaExists(db)
	if checkExistsErr != nil {
		log.Fatalln("DB error: " + checkExistsErr.Error())
	}

	log.Println("check db exists:", schemaExist)

	// create sqlx_test_db database
	if !schemaExist {
		// create schema
		log.Println("schema 建起來")

		_, err := db.Exec("CREATE DATABASE sqlx_test_db CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci")
		if err != nil {
			log.Fatalln("DB error: " + err.Error())
		}

		log.Println("schema 建好了")
	}
	return nil
}

func checkSchemaExists(db *sqlx.DB) (bool, error) {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = 'sqlx_test_db')")
	if err != nil {
		return false, err
	}
	return exists, nil
}

func createTable(db *sqlx.DB) {
	// Check User Schema Exists & Create User
	db.Exec("")

	// Check Post Schema Exists & Create Post
}

func checkTableExists(db *sqlx.DB, tableName string) (bool, error) {
	log.Println("check table exists:", tableName)
	var exists bool

	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'sqlx_test_db' AND TABLE_NAME = ?)", "users")
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	log.Println(tableName, ", table exists:", exists)

	return exists, nil
}

func connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:abc123!@(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	return db, err
}
