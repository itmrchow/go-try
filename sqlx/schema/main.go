package main

import (
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	CREATE_USER_SCHEMA = `
		CREATE TABLE sqlx_test_db.users (
			user_id INT AUTO_INCREMENT PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
			updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			deleted_at DATETIME(6) DEFAULT NULL 
		)
		`

	CREATE_POST_SCHEMA = `
		CREATE TABLE sqlx_test_db.posts (
			post_id INT AUTO_INCREMENT PRIMARY KEY,
			user_id INT NOT NULL,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6),
			updated_at DATETIME(6) DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
			deleted_at DATETIME(6) DEFAULT NULL ,
			FOREIGN KEY (user_id) REFERENCES users(user_id)
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
	tableNames := []string{"users", "posts"}
	for _, tableName := range tableNames {
		if err := createTable(db, tableName); err != nil {
			log.Fatalln("DB error: " + err.Error())
		}
	}

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
func createTable(db *sqlx.DB, tableName string) error {

	// check sqlx_test_db exists
	isUserExists, userExistsErr := checkTableExists(db, tableName)
	if userExistsErr != nil {
		return userExistsErr
	}
	if isUserExists {
		return nil
	}

	log.Println("Create Table:", tableName)

	var execStr string
	switch tableName {
	case "users":
		execStr = CREATE_USER_SCHEMA
	case "posts":
		execStr = CREATE_POST_SCHEMA
	default:
		errorMsg := fmt.Sprintf("table is not definition:%v", tableName)

		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	_, err := db.Exec(execStr)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	log.Println("Create table finish:", tableName)
	return nil
}

func checkTableExists(db *sqlx.DB, tableName string) (bool, error) {
	log.Println("check table exists:", tableName)
	var exists bool

	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'sqlx_test_db' AND TABLE_NAME = ?)", tableName)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	log.Println(tableName, " table exists:", exists)

	return exists, nil
}

func connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:abc123!@(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	return db, err
}
