package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	log.Println("connect db")
	db, err := connectDb()
	if err != nil {
		log.Fatalln("DB error: " + err.Error())
	}
	log.Println("server up")

	// create sqlx_test_db database
	dbExist, checkExistsErr := checkDbExists(db)
	if checkExistsErr != nil {
		log.Fatalln("DB error: " + checkExistsErr.Error())
	}

	if dbExist {
		log.Println("db 在哦")
	} else {
		// create schema
		log.Println("db 建起來")

		log.Println("db 建好了")
	}

	// if err := createSchema(db); err != nil {
	// 	log.Fatalln("DB error: " + err.Error())
	// }

	// start server

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
