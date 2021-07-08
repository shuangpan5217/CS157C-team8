package user

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"CS157C-TEAM8/apis/constants"

	"github.com/go-chi/render"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser = UserPost{}
	var users = []UserPost{}

	querys := r.URL.Query()
	signup := querys.Get("signup")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		render.JSON(w, r, ErrorResponse{
			Message: "Error encounter",
			Error:   err.Error(),
		})
	} else {
		json.Unmarshal(reqBody, &newUser)
		// signup
		if signup == "true" {
			m := make(map[string]interface{})
			iterator := constants.Session.Query("SELECT * FROM "+userTableName+" WHERE username = ? LIMIT 1", newUser.Username).Iter()
			for iterator.MapScan(m) {
				users = append(users, UserPost{
					Username: m["username"].(string),
				})
				m = make(map[string]interface{})
			}
			if len(users) == 0 {
				err = constants.Session.Query(`INSERT INTO `+userTableName+` (username, description, nickname, password, created_time) 
				VALUES(?, ?, ?, ?, ?) IF NOT EXISTS`, newUser.Username, newUser.Description, newUser.Nickname, newUser.Password, time.Now().UTC()).Exec()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					render.JSON(w, r, ErrorResponse{
						Message: "Error encountered",
						Error:   err.Error(),
					})
				} else {
					w.WriteHeader(http.StatusCreated)
					render.JSON(w, r, SuccessResponse{
						Message:       "success",
						MessageStatus: http.StatusCreated,
					})
				}
			} else {
				w.WriteHeader(http.StatusConflict)
				render.JSON(w, r, ErrorResponse{
					Message:       "Error encounter",
					Error:         errors.New("username has been taken").Error(),
					MessageStatus: http.StatusConflict,
				})
			}
		} else { // signin

		}
	}
}
