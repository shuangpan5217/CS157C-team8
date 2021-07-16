package secret

import (
	"CS157C-TEAM8/apis/constants"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"CS157C-TEAM8/apis/user"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func CreateSecretHandler(w http.ResponseWriter, r *http.Request) {
	secretPost := SecretPost{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &secretPost) // new user
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	if secretPost.Username == "" {
		constants.GenerateErrorResponse(w, r, errors.New("username is required in request Body"), http.StatusBadRequest)
		return
	}
	if secretPost.Nickname == "" {
		constants.GenerateErrorResponse(w, r, errors.New("nickname is required in request body"), http.StatusBadRequest)
		return
	}
	if secretPost.Content == "" {
		constants.GenerateErrorResponse(w, r, errors.New("Empty secret content is not allowed."), http.StatusBadRequest)
		return
	}

	if !CheckIfUsernameAndNicknameMatch(secretPost) {
		constants.GenerateErrorResponse(w, r, errors.New("username or nickname is not correct."), http.StatusBadRequest)
		return
	}

	secretID, err := gocql.RandomUUID()
	if err != nil {
		constants.GenerateErrorResponse(w, r, errors.New("Internal Error, please try again."), http.StatusInternalServerError)
		return
	}

	err = CreateSecret(secretID, secretPost, nil)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	GeneratePostSecretSuccessResponse(w, r, "Your secret have been added to the secret box!", http.StatusCreated, secretID)
}

func CreateSecret(secretID gocql.UUID, secretPost SecretPost, createdTime *time.Time) error {
	if createdTime == nil {
		now := time.Now().UTC()
		createdTime = &now
	}
	err := constants.Session.Query(`INSERT INTO `+SecretTableName+` (secret_id, username, nickname, content, created_time) 
	VALUES(?, ?, ?, ?, ?) IF NOT EXISTS`, secretID, secretPost.Username, secretPost.Nickname, secretPost.Content, *createdTime).Exec()

	return err
}

func GetSecretHandler(w http.ResponseWriter, r *http.Request) {
	// check if username is set
	querys := r.URL.Query()
	username := querys.Get("username")
	if username == "" {
		constants.GenerateErrorResponse(w, r, errors.New("username is not set."), http.StatusBadRequest)
		return
	}
	if len(user.GetUserFromDB([]user.UserPost{}, username)) == 0 {
		constants.GenerateErrorResponse(w, r, errors.New("Incorrect username."), http.StatusBadRequest)
		return
	}
	// get ten secret from
	secret, err := GetOneSecretFromDB(username)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	GenerateGetSecretSuccessResponse(w, r, "Successfully get a secret", http.StatusOK, *secret)
}

func GetOneSecretFromDB(username string) (*SecretGet, error) {
	iterator := constants.Session.Query("SELECT * FROM "+SecretTableName+" WHERE token(username) > token(?) LIMIT 5", username).Iter()

	secrets := []SecretGet{}
	m := make(map[string]interface{})
	for iterator.MapScan(m) {
		secrets = append(secrets, SecretGet{
			Username:    m["username"].(string),
			Nickname:    m["nickname"].(string),
			Content:     m["content"].(string),
			SecretID:    cast.ToString(m["secret_id"]),
			CreatedTime: m["created_time"].(time.Time),
		})
		m = make(map[string]interface{})
	}

	iterator = constants.Session.Query("SELECT * FROM "+SecretTableName+" WHERE token(username) < token(?) LIMIT ?", username, 10-len(secrets)).Iter()
	for iterator.MapScan(m) {
		secrets = append(secrets, SecretGet{
			Username:    m["username"].(string),
			Nickname:    m["nickname"].(string),
			Content:     m["content"].(string),
			SecretID:    cast.ToString(m["secret_id"]),
			CreatedTime: m["created_time"].(time.Time),
		})
		m = make(map[string]interface{})
	}

	if len(secrets) == 0 {
		return nil, errors.New("No more secrets, please try again later.")
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(secrets))

	return &secrets[index], nil
}

func DeleteSecretHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	secretDelete := SecretDelete{}
	json.Unmarshal(resp, &secretDelete)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	if secretDelete.SecretID == nil {
		constants.GenerateErrorResponse(w, r, errors.New("secret id is not set."), http.StatusBadRequest)
		return
	}

	if secretDelete.Username == nil {
		constants.GenerateErrorResponse(w, r, errors.New("username is not set."), http.StatusBadRequest)
		return
	}

	secretID := *secretDelete.SecretID
	username := *secretDelete.Username

	_, err = CheckIfSecretExists(secretID, username)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusNotFound)
		return
	}

	err = DeleteSecretFromDB(secretID, username)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	secretUUID, _ := uuid.Parse(secretID)
	GeneratePostSecretSuccessResponse(w, r, "The secret is successfully deleted.", http.StatusOK, gocql.UUID(secretUUID))
}

func CheckIfSecretExists(secretID, username string) (*SecretGet, error) {
	iterator := constants.Session.Query("SELECT * FROM "+SecretTableName+" WHERE username = ? and secret_id = ? LIMIT 1", username, secretID).Iter()
	if iterator.NumRows() == 0 {
		return nil, errors.New("secret doesn't exist.")
	}

	secrets := []SecretGet{}
	m := make(map[string]interface{})
	for iterator.MapScan(m) {
		secrets = append(secrets, SecretGet{
			Username:    m["username"].(string),
			Nickname:    m["nickname"].(string),
			Content:     m["content"].(string),
			SecretID:    cast.ToString(m["secret_id"]),
			CreatedTime: m["created_time"].(time.Time),
		})
		m = make(map[string]interface{})
	}
	return &secrets[0], nil
}

func DeleteSecretFromDB(secretID, username string) error {
	err := constants.Session.Query("DELETE from "+SecretTableName+" WHERE username = ? and secret_id = ?", username, secretID).Exec()
	if err != nil {
		return err
	}
	return nil
}

func CheckIfUsernameAndNicknameMatch(secretPost SecretPost) bool {
	users := user.GetUserFromDB([]user.UserPost{}, secretPost.Username)
	if len(users) == 0 {
		return false
	}
	if secretPost.Nickname != users[0].Nickname {
		return false
	}
	return true
}
