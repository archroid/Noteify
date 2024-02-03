package database

import (
	"archroid/noteify/models"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Init() error {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}

	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS users (
		ID INTEGER PRIMARY KEY,
		username TEXT,
		password TEXT,
		created_at INTEGER,
		token TEXT
	)`)
	if err != nil {
		return err
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
		return err
	}

	return nil

}

func AddUser(user models.User) error {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	database.Exec("INSERT INTO users (username, password, created_at, token, id) VALUES (?, ?, ?, ?,?)", user.Username, user.Password, user.CREATED_AT, user.Token, user.ID)
	return nil
}

func GetUser(username string) (models.User, error) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return models.User{}, err
	}
	rows, _ := database.Query("SELECT * FROM users WHERE username = ?", username)
	var user models.User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.CREATED_AT, &user.Token)
	}
	return user, nil
}

func DeleteUser(username string) error {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	database.Exec("DELETE FROM users WHERE username = ?", username)
	return nil
}


func AddPost(post models.Post) error {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	database.Exec("INSERT INTO Posts (PostTitle, PostDate, Deleted, OwnerID) VALUES (?, ?, ?, ?)", post.PostTitle, post.PostDate, post.Deleted, post.OwnerID)
	return nil
}

func GetPosts() ([]models.Post, error) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return []models.Post{}, err
	}
	rows, _ := database.Query("SELECT * FROM Posts")
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(postID int) (models.Post, error) {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return models.Post{}, err
	}
	rows, _ := database.Query("SELECT * FROM Posts WHERE PostID = ?", postID)
	var post models.Post
	for rows.Next() {
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
	}
	return post, nil
}

func DeletePost(postID int) error {
	database, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		return err
	}
	database.Exec("DELETE FROM Posts WHERE PostID = ?", postID)
	return nil
}

