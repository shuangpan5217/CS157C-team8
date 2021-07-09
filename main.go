package main

import (
	"fmt"
	"log"
	"net/http"

	constants "CS157C-TEAM8/apis/constants"
	"CS157C-TEAM8/apis/user"

	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "pan"
	constants.Session, err = cluster.CreateSession()

	defer constants.Session.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/login", user.CreateUser).Methods("POST")
	router.HandleFunc("/updateuser", user.UpdateUser).Methods("Patch")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("connecting to localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
