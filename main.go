package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"CS157C-TEAM8/apis/constants"
	"CS157C-TEAM8/apis/savedsecrets"
	"CS157C-TEAM8/apis/secret"
	"CS157C-TEAM8/apis/user"

	"github.com/go-chi/render"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	// set default host
	// will be overrided by docker env variable when run 'docker run' command
	if os.Getenv(constants.CASSANDRA_URL) == "" {
		os.Setenv(constants.CASSANDRA_URL, "127.0.0.1:9042")
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, map[string]string{
		"Message": "Welcome to use CS157C-TEAM8 api services",
	})
}

func main() {
	var err error

	cluster := constants.InitilizeCluster()
	constants.Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println(" => cassandra well initialized")

	defer constants.Session.Close()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)

	router.HandleFunc("/login", user.CreateUserHandler).Methods("POST")
	router.HandleFunc("/updateuser", user.UpdateUserHandler).Methods("Patch")

	router.HandleFunc("/secret", secret.CreateSecretHandler).Methods("POST")
	router.HandleFunc("/secret", secret.GetSecretHandler).Methods("GET")
	router.HandleFunc("/secret", secret.DeleteSecretHandler).Methods("DELETE")

	router.HandleFunc("/savedsecret", savedsecrets.SaveSecretHandler).Methods("POST")
	router.HandleFunc("/savedsecret", savedsecrets.GetAllFavoriteSecretsHandler).Methods("GET")
	router.HandleFunc("/savedsecret", savedsecrets.RemoveSavedSecretHandler).Methods("DELETE")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println(" => connecting to localhost:4999")
	log.Fatal(http.ListenAndServe(":4999", handlers.CORS(headers, methods, origins)(router)))
}
