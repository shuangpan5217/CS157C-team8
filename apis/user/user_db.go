package user

import (
	"net/http"
	"time"

	constants "CS157C-TEAM8/apis/constants"

	"github.com/go-chi/render"
)

var userTableName = "user"

type UserPost struct {
	Username    string     `json:"username"`
	Description string     `json:"description"`
	Nickname    string     `json:"nickname"`
	Password    string     `json:"password"`
	CreatedTime *time.Time `json:"created_time"`
}

type UserUpdate struct {
	Username    string  `json:"username"`
	Description *string `json:"description,omitempty"`
	Nickname    *string `json:"nickname,omitempty"`
}

func GenerateUserSuccessResponse(w http.ResponseWriter, r *http.Request, successMessage string, statusCode int, user UserPost) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, constants.SuccessResponse{
		Message:        "success",
		SuccessMessage: successMessage,
		StatusCode:     statusCode,
		Body: map[string]string{
			"usernaame":   user.Username,
			"nickname":    user.Nickname,
			"description": user.Description,
		},
	})
}
