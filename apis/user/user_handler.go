package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"CS157C-TEAM8/apis/constants"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser = UserPost{} // used to fetch the user from Cassandra
	var users = []UserPost{} // used to check if username exists or not

	querys := r.URL.Query()
	signup := querys.Get("signup")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &newUser) // new user
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	if newUser.Username == "" {
		constants.GenerateErrorResponse(w, r, errors.New("username is required in request Body"), http.StatusBadRequest)
		return
	}
	if newUser.Password == "" {
		constants.GenerateErrorResponse(w, r, errors.New("password is required in request Body"), http.StatusBadRequest)
		return
	}

	users = GetUserFromDB(users, newUser.Username)

	// signup
	if signup == "true" {
		// if useranme is not taken
		if len(users) == 0 {
			if newUser.Nickname == "" {
				newUser.Nickname = GenerateRandomNickname(10)
			}
			err = constants.Session.Query(`INSERT INTO `+userTableName+` (username, description, nickname, password, created_time) 
			VALUES(?, ?, ?, ?, ?) IF NOT EXISTS`, newUser.Username, newUser.Description, newUser.Nickname, newUser.Password, time.Now().UTC()).Exec()
			if err != nil {
				constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
			} else {
				GenerateUserSuccessResponse(w, r, "Successfully create an user.", http.StatusCreated, UserPost{
					Username:    newUser.Username,
					Nickname:    newUser.Nickname,
					Description: newUser.Description,
				})
			}
		} else {
			constants.GenerateErrorResponse(w, r, errors.New("username has been taken."), http.StatusConflict)
		}
	} else if signup != "true" && signup != "" {
		constants.GenerateErrorResponse(w, r, errors.New("Please assign correct query string."), http.StatusBadRequest)
	} else {
		// signin
		// if there is no username in the database
		if len(users) == 0 {
			constants.GenerateErrorResponse(w, r, errors.New("username or password is not correct."), http.StatusUnauthorized)
			return
		}
		// check password
		if users[0].Password != newUser.Password {
			constants.GenerateErrorResponse(w, r, errors.New("username or password is not correct."), http.StatusUnauthorized)
		} else {
			GenerateUserSuccessResponse(w, r, "Successfully log in.", http.StatusOK, UserPost{
				Username:    users[0].Username,
				Nickname:    users[0].Nickname,
				Description: users[0].Description,
			})
		}
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user = UserUpdate{} // used to update the user from Cassandra
	var users = []UserPost{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}
	users = GetUserFromDB(users, user.Username)
	if len(users) == 0 {
		constants.GenerateErrorResponse(w, r, errors.New("user doesn't exist."), http.StatusUnauthorized)
		return
	}
	if user.Nickname != nil && user.Description != nil {
		err = constants.Session.Query("UPDATE "+userTableName+` set description = ?, nickname = ? WHERE username = ?`, user.Description, user.Nickname, user.Username).Exec()
	} else if user.Nickname != nil {
		err = constants.Session.Query("UPDATE "+userTableName+` set nickname = ? WHERE username = ?`, user.Nickname, user.Username).Exec()
	} else if user.Description != nil {
		err = constants.Session.Query("UPDATE "+userTableName+` set description = ? WHERE username = ?`, user.Description, user.Username).Exec()
	}

	if err != nil {
		constants.GenerateErrorResponse(w, r, err, http.StatusInternalServerError)
		return
	}

	users = GetUserFromDB([]UserPost{}, user.Username)
	GenerateUserSuccessResponse(w, r, "Successfully updated.", http.StatusOK, UserPost{
		Username:    users[0].Username,
		Nickname:    users[0].Nickname,
		Description: users[0].Description,
	})

}

func GetUserFromDB(users []UserPost, username string) []UserPost {
	m := make(map[string]interface{})
	// get user in database

	iterator := constants.Session.Query(`SELECT * FROM `+userTableName+` WHERE username = ? LIMIT 1`, username).Iter()
	for iterator.MapScan(m) {
		users = append(users, UserPost{
			Username:    m["username"].(string),
			Password:    m["password"].(string),
			Nickname:    m["nickname"].(string),
			Description: m["description"].(string),
		})
		m = make(map[string]interface{})
	}
	iterator.Close()
	return users
}

func GenerateRandomNickname(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
