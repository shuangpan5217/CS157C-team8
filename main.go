package main

import (
	"fmt"
	"log"
	"net/http"

	constants "CS157C-TEAM8-PROJECT-API/constants"
	"CS157C-TEAM8-PROJECT-API/user"

	"github.com/gocql/gocql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "pan"
	constants.Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("cassandra well initialized")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/createuser", user.CreateUser).Methods("POST")
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(headers, methods, origins)(router)))
}
