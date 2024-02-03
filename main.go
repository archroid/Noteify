package main

import (
	"archroid/noteify/database"
	"archroid/noteify/models"
	"encoding/json"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var localip string

func main() {
	database.Init()

	//find device local ip
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			localip = ipv4.String()
		}
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/posts", GetPostsHandler)

	log.Info("Server started: http://" + localip + ":8090")
	log.Error(http.ListenAndServe(":8090", r))
	// go func() {

	// }()
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
