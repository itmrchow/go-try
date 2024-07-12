package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
	// insertPosts(db)

	// updatePost(db)
	// updatePosts(db)

	// deletePost(db)

	getPost(db)

}

type Post struct {
	PostId    int       `db:"post_id"`
	UserId    int       `db:"user_id"`
	Title     string    `db:"title"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}

func getPost(db *sqlx.DB) {
	println("Get post")

	post := Post{}
	postId := 9
	sqlStr := `SELECT post_id , user_id , title , content , created_at , updated_at , deleted_at FROM posts WHERE post_id = ?`

	err := db.Get(&post, sqlStr, postId)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	if err != nil {
		log.Fatalln("Get error:", err.Error())
		return
	}

	fmt.Printf("%+v\n", post)
	print("UnixNano:")
	println(post.DeletedAt.UnixNano())

	return
}

func deletePost(db *sqlx.DB) {
	log.Println("Delete post")

	currentTime := time.Now()
	print("Unix:")
	println(currentTime.UnixMicro())

	arg := map[string]interface{}{
		"deleted_at": currentTime,
		"post_ids":   []int{9, 11, 13, 15},
	}

	deleteStr := "UPDATE `sqlx_test_db`.`posts` SET `deleted_at` = :deleted_at WHERE `post_id` IN (:post_ids)"
	query, args, err := sqlx.Named(deleteStr, arg)
	if err != nil {
		log.Fatalln("Delete error:", err.Error())
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		log.Fatalln("Delete error:", err.Error())
	}

	query = db.Rebind(query)

	_, err = db.Exec(query, args...)
	if err != nil {
		log.Fatalln("Delete error:", err.Error())
	}
}

func updatePost(db *sqlx.DB) {
	log.Println("Update Post")

	updateStr := "UPDATE `sqlx_test_db`.`posts` SET `title` = :title, `content` = :content WHERE `post_id` = :post_id"
	_, err := db.NamedExec(updateStr, &Post{PostId: 2, Title: "PostTitleUpdated", Content: "ContentUpdated"})

	if err != nil {
		log.Fatalln("Update error:", err.Error())
	}
}

func updatePosts(db *sqlx.DB) {
	log.Println("Update Posts")

	updateStr := "UPDATE `sqlx_test_db`.`posts` SET `content` = :content WHERE `title` LIKE :msg"
	_, err := db.NamedExec(updateStr, map[string]interface{}{"msg": "%AAA", "content": "update:AAA"})

	if err != nil {
		log.Fatalln("Update error:", err.Error())
	}
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

	for i := 0; i < 30000; i++ {
		posts = append(posts, Post{UserId: 1, Title: fmt.Sprintf("PostTitle%d", i), Content: fmt.Sprintf("Content%d", i)})
	}

	for i := 0; i < len(posts); i += 100 {
		insertStr := "INSERT INTO `sqlx_test_db`.`posts` (`user_id`, `title`, `content` ) VALUES (:user_id, :title, :content)"
		_, err := db.NamedExec(insertStr, posts[i:i+100])
		if err != nil {
			log.Fatalln("Insert error:", err.Error())
		}
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
	db, err := sqlx.Connect("mysql", "root:abc123!@(localhost:3306)/sqlx_test_db?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, err
}
