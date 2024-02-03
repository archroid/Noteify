package main

import (
	"archroid/noteify/database"
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

	r.HandleFunc("/api/user/add", AddUserHanlder).Methods("POST")
	r.HandleFunc("/api/user/login", LoginUserHanlder).Methods("POST")

	r.HandleFunc("/api/post/add", AddPostHandler).Methods("POST")
	r.HandleFunc("/api/posts", GetPostsHandler).Methods("GET")
	r.HandleFunc("/api/post/{username}", GetUserPostHandler).Methods("GET")

	log.Info("Server started: http://" + localip + ":8090")
	log.Error(http.ListenAndServe(":8090", r))

}
