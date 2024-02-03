package main

import (
	"archroid/noteify/database"
	"archroid/noteify/models"
	"archroid/noteify/util"
	"encoding/json"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	posts, err := database.GetPosts()
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddUserHanlder(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user := models.User{
		Username:   username,
		Password:   password,
		Token:      util.GenerateToken(username),
		CREATED_AT: int(time.Now().Unix()),
	}

	_, err := database.GetUser(username)
	if err != nil {
		database.AddUser(user)
		w.WriteHeader(http.StatusOK)
	} else {
		log.Error(err)
		w.WriteHeader(http.StatusConflict)
	}
}
