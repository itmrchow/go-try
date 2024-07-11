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

	createUser(db)

	// createPost(db)

}

type User struct {
	UserId    int    `db:"user_id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func createUser(db *sqlx.DB) {

	insertStr := "INSERT INTO `sqlx_test_db`.`users` (`username`, `email`, `password` ) VALUES (:username, :email, :password)"
	_, err := db.NamedExec(insertStr, &User{Username: "Jojo", Email: "Jojo@gmail.com", Password: "pwd"})

	if err != nil {
		log.Fatalln("Insert error:", err.Error())
	}
}

func connectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:abc123!@(localhost:3306)/")
	if err != nil {
		return nil, err
	}

	return db, err
}
