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

	// getPost(db)
	// findPost(db)

	findUserAndPost(db)
}

type page struct {
	id    int
	limit int
	flip  string
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

type User struct {
	UserID    int        `db:"user_id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	Posts []Post `db:"-"`
}

type UserAndLatestPost struct {
	UserID    int        `db:"user_id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	Post
}

func findUserAndPost(db *sqlx.DB) {
	userId := 2

	startV1 := time.Now()
	findUserAndPostV1(db, userId)
	elapsedV1 := time.Since(startV1)

	fmt.Printf("Run time v1: %v\n", elapsedV1)

	startV2 := time.Now()
	findUserAndPostV2(db, userId)
	elapsedV2 := time.Since(startV2)

	fmt.Printf("Run time v1: %v\n", elapsedV2)
}

func findUserAndPostV1(db *sqlx.DB, userId int) {
	var maxPostId int
	// v1 : 先查max post , 再組起來
	// v1-1 query max post id
	queryMaxPost := `
		SELECT Max(posts.post_id)  
		FROM posts 
		WHERE posts.user_id = ? AND posts.deleted_at IS NULL;
	`
	queryMaxIdErr := db.Get(&maxPostId, queryMaxPost, userId)
	if queryMaxIdErr != nil {
		log.Fatalln("findUserAndPost error:", queryMaxIdErr.Error())
		return
	}
	fmt.Printf("%+v\n", maxPostId)

	// v1-2 query max post
	post := Post{}
	db.Get(&post, `
		SELECT post_id , user_id , title , content , created_at , updated_at
		FROM posts
		WHERE post_id = ?
	`, maxPostId)

	fmt.Printf("%+v\n", post)

	// v1-2 query User
	user := User{}
	db.Get(&user,
		`
		SELECT user_id, username, email, password, created_at, updated_at
		FROM users
		WHERE user_id = ?
	`, userId)

	user.Posts = append(user.Posts, post)

	fmt.Printf("%+v\n", user)
}

func findUserAndPostV2(db *sqlx.DB, userId int) {
	// v2 : 寫在同一個query
	var maxPostId int
	// v2-1 query max post id
	queryMaxPost := `
		SELECT Max(posts.post_id)  
		FROM posts 
		WHERE posts.user_id = ? AND posts.deleted_at IS NULL;
	`
	queryMaxIdErr := db.Get(&maxPostId, queryMaxPost, userId)
	if queryMaxIdErr != nil {
		log.Fatalln("findUserAndPost error:", queryMaxIdErr.Error())
		return
	}
	fmt.Printf("%+v\n", maxPostId)

	user := UserAndLatestPost{}
	queryUser := `
		SELECT 
			users.user_id,users.username,users.email,users.password,users.created_at,users.updated_at,users.deleted_at,
		 	posts.post_id,posts.user_id,posts.title,posts.content,posts.created_at,posts.updated_at,posts.deleted_at
		FROM users
		LEFT JOIN posts ON 
			users.user_id = posts.user_id
			AND
			posts.post_id = ?
		WHERE users.user_id = ?
	`
	queryUserErr := db.Get(&user, queryUser, maxPostId, userId)
	if queryUserErr != nil {
		log.Fatalln("findUserAndPost error:", queryUserErr.Error())
		return
	}
	fmt.Printf("%+v\n", user)
}

// Query Post , delete no show , 分頁 , 指定User ,  指定比數後的 , 新的在後
func findPost(db *sqlx.DB) {

	page := page{
		id:    13,
		limit: 20,
		flip:  "next",
	}

	sort := "DESC"
	sqlStr := `
	SELECT post_id , user_id , title , content , created_at , updated_at
	FROM posts
	WHERE
		user_id = :user_id
		AND
		deleted_at IS NULL
		ORDER BY post_id %s
		LIMIT :id , :limit
	`

	if sort == "DESC" {
		sqlStr = fmt.Sprintf(sqlStr, sort)
	} else {
		sqlStr = fmt.Sprintf(sqlStr, "ASC")
	}

	arg := map[string]interface{}{
		"user_id": 2,
		"sort":    sort,
		"id":      page.id,
		"limit":   page.limit,
	}

	posts := []Post{}

	// 塞值
	query, args, namedErr := sqlx.Named(sqlStr, arg)
	if namedErr != nil {
		log.Fatalln("findPost error:", namedErr.Error())
		return
	}

	fmt.Printf("%+v\n", query)
	fmt.Printf("%+v\n", args)

	// 查詢
	selectErr := db.Select(&posts, query, args...)
	if selectErr != nil {
		log.Fatalln("findPost error:", selectErr.Error())
		return
	}

	for _, post := range posts {
		fmt.Printf("%+v\n", post)
	}

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

	for i := 0; i < 100000; i++ {
		posts = append(posts, Post{UserId: 2, Title: fmt.Sprintf("PostTitle%d", i), Content: fmt.Sprintf("Content%d", i)})
	}

	for i := 0; i < len(posts); i += 100 {
		insertStr := "INSERT INTO `sqlx_test_db`.`posts` (`user_id`, `title`, `content` ) VALUES (:user_id, :title, :content)"
		_, err := db.NamedExec(insertStr, posts[i:i+100])
		if err != nil {
			log.Fatalln("Insert error:", err.Error())
		}
	}
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
