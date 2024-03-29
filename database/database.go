package database

import (
	"archroid/noteify/models"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Init() error {
	var err error
	database, err = sql.Open("sqlite3", "./data.db")
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
	_, err := database.Exec("INSERT INTO users (username, password, created_at, token) VALUES (?, ?, ?, ?)", user.Username, user.Password, user.CREATED_AT, user.Token)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(username string) (models.User, error) {
	rows, err := database.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.CREATED_AT, &user.Token)
		if err != nil {
			return models.User{}, errors.New("user not found")
		}
	} else {
		return models.User{}, errors.New("user not found")
	}

	return user, nil
}

func DeleteUser(username string) error {
	_, err := database.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		return err
	}
	return nil
}

func AddPost(post models.Post) error {
	_, err := database.Exec("INSERT INTO Posts (PostTitle, PostDate, Deleted, OwnerID) VALUES (?, ?, ?, ?)", post.PostTitle, post.PostDate, post.Deleted, post.OwnerID)
	if err != nil {
		return err
	}
	return nil
}

func GetPosts() ([]models.Post, error) {
	var posts []models.Post

	rows, err := database.Query("SELECT * FROM Posts")
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
		posts = append(posts, post)
	}
	return posts, nil
}

func GetUserPosts(username string) ([]models.Post, error) {
	var posts []models.Post
	user, err := GetUser(username)
	if err != nil {
		return posts, err
	}
	rows, err := database.Query("SELECT * FROM Posts WHERE OwnerID = ?", user.ID)
	if err != nil {
		return posts, err
	}
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(postID int) (models.Post, error) {
	var post models.Post

	rows, err := database.Query("SELECT * FROM Posts WHERE PostID = ?", postID)
	if err != nil {
		return post, err
	}
	for rows.Next() {
		rows.Scan(&post.PostID, &post.PostTitle, &post.PostDate, &post.Deleted, &post.OwnerID)
	}
	return post, nil
}

func DeletePost(postID int) error {
	_, err := database.Exec("DELETE FROM Posts WHERE PostID = ?", postID)
	if err != nil {
		return err
	}
	return nil
}
