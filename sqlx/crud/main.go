package main

import (
	"fmt"
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

	// insertUser(db)

	// insertPost(db)
	insertPosts(db)

	// updatePost(db)

}

type Post struct {
	PostId    int    `db:"post_id"`
	UserId    int    `db:"user_id"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
	DeletedAt int64  `db:"deleted_at"`
}

func insertPost(db *sqlx.DB) {
	log.Println("Create Post")

	insertStr := "INSERT INTO `sqlx_test_db`.`posts` (`user_id`, `title`, `content` ) VALUES (:user_id, :title, :content)"
	_, err := db.NamedExec(insertStr, &Post{UserId: 1, Title: "PostTitle0", Content: "Content0"})

	if err != nil {
		log.Fatalln("Insert error:", err.Error())
	}
}

func insertPosts(db *sqlx.DB) {
	log.Println("Create Post")

	posts := []Post{}

	for i := 0; i < 100; i++ {
		posts = append(posts, Post{UserId: 1, Title: fmt.Sprintf("PostTitle%d", i), Content: fmt.Sprintf("Content%d", i)})
	}

	insertStr := "INSERT INTO `sqlx_test_db`.`posts` (`user_id`, `title`, `content` ) VALUES (:user_id, :title, :content)"
	_, err := db.NamedExec(insertStr, posts)
	if err != nil {
		log.Fatalln("Insert error:", err.Error())
	}
}

type User struct {
	UserId    int    `db:"user_id"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt int64  `db:"created_at"`
	UpdatedAt int64  `db:"updated_at"`
}

func insertUser(db *sqlx.DB) {
	log.Println("Create user")

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
