package database

import (
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
		username TEXT,
		password TEXT,
		ID INTEGER PRIMARY KEY,
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

// statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT) ")
// statement.Exec()
// statement, _ = database.Prepare("INSERT INTO fuckers (firstname,lastname) VALUES (?,?)")
// statement.Exec("Nic", "Reboy")

// rows, _ := database.Query("SELECT * FROM fuckers")
// var id int
// var firstname string
// var lastname string
// for rows.Next() {
// 	rows.Scan(&id, &firstname, &lastname)
// 	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
// }
