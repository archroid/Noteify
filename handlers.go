package main

import (
	"archroid/noteify/database"
	"archroid/noteify/models"
	"archroid/noteify/util"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	ownerID, err := strconv.Atoi(r.PostFormValue("ownerid"))
	if err != nil {
		log.Error(err)
	}

	post := models.Post{
		PostTitle: r.PostFormValue("title"),
		PostDate:  r.PostFormValue("date"),
		Deleted:   0,
		OwnerID:   ownerID,
	}

	err = database.AddPost(post)
	if err != nil {
		log.Error(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post added successfully"))
}

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

func GetUserPostHandler(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	username := mux.Vars(r)["username"]
	posts, err := database.GetUserPosts(username)
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

	founduser, _ := database.GetUser(username)

	if founduser.Username == user.Username {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("User already exists"))
	} else {
		err := database.AddUser(user)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("User added successfully"))
		}
	}

}

func LoginUserHanlder(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	user, err := database.GetUser(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	if user.Password == password {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid credentials"))
	}

}
