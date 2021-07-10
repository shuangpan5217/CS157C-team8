package constants

import (
	"fmt"
	"log"
	"os"

	"github.com/gocql/gocql"
)

const (
	CASSANDRA_URL          = "CASSANDRA_URL"
	KeySpaceName           = "pan"
	createKeySpace         = "CREATE KEYSPACE IF NOT EXISTS " + KeySpaceName + " WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }"
	createUserTable        = "CREATE TABLE IF NOT EXISTS " + KeySpaceName + ".user (username text, password text, nickname text, description text, created_time timestamp, primary key(username))"
	createSecretTable      = "create table IF NOT EXISTS " + KeySpaceName + ".secret (secret_id uuid, username text, nickname text, content text, created_time timestamp, primary key(secret_id))"
	createSavedSecretTable = "create table IF NOT EXISTS " + KeySpaceName + ".saved_secret(secret_id uuid, username text, content text, nickname text, primary key(secret_id, username))"
	createCommentTable     = "create table IF NOT EXISTS " + KeySpaceName + ".comment (comment_id uuid, secret_id uuid, created_time timestamp, comment text, nickname text, primary key(comment_id))"
)

func InitilizeCluster() (cluster *gocql.ClusterConfig) {
	// look up for env variable
	host, exist := os.LookupEnv("CASSANDRA_URL")
	if !exist {
		log.Fatal("Host doesn't exist.")
	}

	// create cluster
	cluster = gocql.NewCluster(host)
	fmt.Println(" => connect to cluster host: " + host)

	// create keyspace if it doesn't exist
	keySpace := CreateKeySpace(cluster)
	cluster.Keyspace = keySpace

	return
}

func CreateKeySpace(cluster *gocql.ClusterConfig) (keySpace string) {
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	session.Query(createKeySpace).Exec()

	createTable(session, createUserTable)
	createTable(session, createSecretTable)
	createTable(session, createSecretTable)
	createTable(session, createCommentTable)

	return KeySpaceName
}

func createTable(session *gocql.Session, tableName string) {
	session.Query(tableName).Exec()
}
