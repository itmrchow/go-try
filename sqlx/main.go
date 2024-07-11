package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := connectDb()
	if err != nil {
		log.Fatalln("DB error: " + err.Error())
	}
	log.Println("connected db")

	if err := createDb(db); err != nil {
		log.Fatalln("DB error: " + err.Error())
	}

}

func createDb(db *sqlx.DB) error {
	// check sqlx_test_db exists
	dbExist, checkExistsErr := checkDbExists(db)
	if checkExistsErr != nil {
		log.Fatalln("DB error: " + checkExistsErr.Error())
	}

	log.Println("check db exists:", dbExist)

	// create sqlx_test_db database
	if !dbExist {
		// create db
		log.Println("db 建起來")

		_, err := db.Exec("CREATE DATABASE sqlx_test_db CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci")
		if err != nil {
			log.Fatalln("DB error: " + err.Error())
		}

		log.Println("db 建好了")
	}
	return nil
}

func checkDbExists(db *sqlx.DB) (bool, error) {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = 'sqlx_test_db')")
	if err != nil {
		return false, err
	}
	return exists, nil
}

func connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:abc123!@(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	return db, err
}
