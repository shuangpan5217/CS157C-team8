package savedsecrets

import (
	"CS157C-TEAM8/apis/constants"
	"CS157C-TEAM8/apis/secret"
	"CS157C-TEAM8/apis/user"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

func SaveSecretHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	savedSecret := SavedSecretPost{}

	err = json.Unmarshal(resp, &savedSecret)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	// will not check this because it may be deleted in the process of adding to favorite list
	secretID := savedSecret.SecretID

	// check if matches
	secretOwner := savedSecret.SecretOwner // the username who created the secret
	nickname := savedSecret.Nickname
	matches := secret.CheckIfUsernameAndNicknameMatch(secret.SecretPost{Username: secretOwner, Nickname: nickname})
	if !matches {
		constants.GenerateErrorResponse(w, r, errors.New("secret owner's username or nickname is not correct."), http.StatusBadRequest)
		return
	}

	content := savedSecret.Content // not empty
	if content == "" {
		constants.GenerateErrorResponse(w, r, errors.New("Empty secret content is not allowed."), http.StatusBadRequest)
		return
	}
	// will not check these two. Doesn't matter
	createdTime := savedSecret.CreatedTime
	if cast.ToString(createdTime) == "" {
		constants.GenerateErrorResponse(w, r, errors.New("created_tiem is not set or empty."), http.StatusBadRequest)
		return
	}

	username := savedSecret.Username // the username who saved the secret, added to favorite list
	users := user.GetUserFromDB([]user.UserPost{}, username)
	if len(users) == 0 {
		constants.GenerateErrorResponse(w, r, errors.New("user doesn't exist."), http.StatusBadRequest)
		return
	}

	err = AddToFavoriteList(savedSecret)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	err = secret.DeleteSecretFromDB(secretID, secretOwner)
	if err != nil {
		constants.GenerateErrorResponse(w, r, errors.New("Internal Error, please try again."), http.StatusInternalServerError)
		return
	}

	GeneratePostSavedSecretSuccessResponse(w, r, "Successfully added to your favorite list.", http.StatusOK, savedSecret)
}

func AddToFavoriteList(savedSecret SavedSecretPost) error {
	err := constants.Session.Query("INSERT INTO "+SavedSecretsTableName+"(secret_id, secret_owner, username, content, nickname, created_time) VALUES(?, ?, ?, ?, ?, ?) IF NOT EXISTS", savedSecret.SecretID, savedSecret.SecretOwner, savedSecret.Username, savedSecret.Content, savedSecret.Nickname, savedSecret.CreatedTime).Exec()
	return err
}

func GetAllFavoriteSecretsHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	savedSecret := SavedSecretPost{}

	err = json.Unmarshal(resp, &savedSecret)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	users := user.GetUserFromDB([]user.UserPost{}, savedSecret.Username)
	if len(users) == 0 {
		constants.GenerateErrorResponse(w, r, errors.New("username doesn't exist"), http.StatusNotFound)
		return
	}

	savedSecrets := []SavedSecretPost{}
	iterator := constants.Session.Query("SELECT * FROM "+SavedSecretsTableName+" WHERE username = ? ALLOW FILTERING", savedSecret.Username).Iter()

	m := make(map[string]interface{})
	for iterator.MapScan(m) {
		savedSecrets = append(savedSecrets, SavedSecretPost{
			Username:    m["username"].(string),
			Nickname:    m["nickname"].(string),
			Content:     m["content"].(string),
			SecretID:    cast.ToString(m["secret_id"]),
			CreatedTime: m["created_time"].(time.Time),
		})
		m = make(map[string]interface{})
	}
	iterator.Close()

	GenerateGetSavedSecretSuccessResponse(w, r, "All saved secrets", http.StatusOK, savedSecrets)
}

func RemoveSavedSecretHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	savedSecret := SavedSecretPost{}

	err = json.Unmarshal(resp, &savedSecret)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	secretPosts := CheckIfSecretExistsInFavoriteList(savedSecret)
	if len(secretPosts) == 0 {
		constants.GenerateErrorResponse(w, r, errors.New("Saved secret Not found"), http.StatusNotFound)
		return
	}

	querys := r.URL.Query()
	throwback := querys.Get("throwback")
	if strings.ToLower(throwback) == "true" {
		secretUUID, _ := uuid.Parse(secretPosts[0].SecretID)
		err = secret.CreateSecret(gocql.UUID(secretUUID), secretPosts[0], &secretPosts[0].CreatedTime)
		if err != nil {
			constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
			return
		}
	}

	err = RemoveSavedSecretFromDB(savedSecret)
	if err != nil {
		constants.GenerateErrorResponse(w, r, errors.New("Fail removing your saved secret, please try again."), http.StatusInternalServerError)
		return
	}

	GenerateRemoveSavedSecretSuccessResponse(w, r, "Successfully removed from your favorite list.", http.StatusOK, savedSecret.SecretID)
}

func RemoveSavedSecretFromDB(savedSecret SavedSecretPost) error {
	err := constants.Session.Query("DELETE FROM "+SavedSecretsTableName+" WHERE username = ? and secret_id = ?", savedSecret.Username, savedSecret.SecretID).Exec()
	return err
}

func CheckIfSecretExistsInFavoriteList(savedSecret SavedSecretPost) []secret.SecretPost {
	iterator := constants.Session.Query("SELECT * FROM "+SavedSecretsTableName+" WHERE username = ? and secret_id = ?", savedSecret.Username, savedSecret.SecretID).Iter()
	secretPosts := []secret.SecretPost{}
	m := make(map[string]interface{})
	for iterator.MapScan(m) {
		secretPosts = append(secretPosts, secret.SecretPost{
			Username:    m["secret_owner"].(string),
			Nickname:    m["nickname"].(string),
			Content:     m["content"].(string),
			SecretID:    cast.ToString(m["secret_id"]),
			CreatedTime: m["created_time"].(time.Time),
		})
		m = make(map[string]interface{})
	}
	return secretPosts
}
