package database

import (
	"archroid/noteify/models"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func Init() {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}

	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS users (
		ID INTEGER PRIMARY KEY,
		username TEXT,
		password TEXT,
		created_at INTEGER,
		token TEXT
	)`)
	if err != nil {
		logrus.Error(err)
	}

	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS Posts (
		PostID INTEGER PRIMARY KEY,
		PostTitle TEXT,
		PostDate TEXT,
		Deleted INTEGER,
		OwnerID INTEGER,
		FOREIGN KEY(OwnerID) REFERENCES users(id)
		)`)
	if err != nil {
		logrus.Error(err)
	}

}

func AddUser(user models.User) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	database.Exec("INSERT INTO users (username, password, created_at, token, id) VALUES (?, ?, ?, ?,?)", user.Username, user.Password, user.CREATED_AT, user.Token, user.ID)
}

func GetUser(username string) models.User {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	rows, _ := database.Query("SELECT * FROM users WHERE username = ?", username)
	var user models.User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.CREATED_AT, &user.Token)
	}
	return user
}

func DeleteUser(username string) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	database.Exec("DELETE FROM users WHERE username = ?", username)
}

func AddPost(post models.Post) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	database.Exec("INSERT INTO Posts (PostTitle, PostDate, Deleted, OwnerID) VALUES (?, ?, ?, ?)", post.PostTitle, post.PostDate, post.Deleted, post.OwnerID)
}

func GetPosts() []models.Post {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	rows, _ := database.Query("SELECT * FROM Posts")
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
		posts = append(posts, post)
	}
	return posts
}

func GetPost(postID int) models.Post {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	rows, _ := database.Query("SELECT * FROM Posts WHERE PostID = ?", postID)
	var post models.Post
	for rows.Next() {
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
	}
	return post
}

func DeletePost(postID int) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		logrus.Error(err)
	}
	database.Exec("DELETE FROM Posts WHERE PostID = ?", postID)
}
