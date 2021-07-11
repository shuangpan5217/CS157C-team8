package secret

import (
	"CS157C-TEAM8/apis/constants"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/gocql/gocql"
	"github.com/spf13/cast"
)

var SecretTableName = "secret"

type SecretPost struct {
	SecretID    string     `json:"secret_id"`
	Content     string     `json:"content"`
	CreatedTime *time.Time `json:"created_time"`
	Nickname    string     `json:"nickname"`
	Username    string     `json:"username"`
}

type SecretGet struct {
	SecretID    string     `json:"secret_id"`
	Content     string     `json:"content"`
	CreatedTime *time.Time `json:"created_time"`
	Nickname    string     `json:"nickname"`
}

func GenerateGetSecretSuccessResponse(w http.ResponseWriter, r *http.Request, successMessage string, statusCode int, secret SecretGet) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, constants.SuccessResponse{
		Message:        "success",
		SuccessMessage: successMessage,
		StatusCode:     statusCode,
		Body: map[string]string{
			"nickname":     secret.Nickname,
			"content":      secret.Content,
			"secret_id":    secret.SecretID,
			"created_time": cast.ToString(secret.CreatedTime),
		},
	})
}

func GeneratePostSecretSuccessResponse(w http.ResponseWriter, r *http.Request, successMessage string, statusCode int, secretID gocql.UUID) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, constants.SuccessResponse{
		Message:        "success",
		SuccessMessage: successMessage,
		StatusCode:     statusCode,
		Body: map[string]string{
			"secret_id": cast.ToString(secretID),
		},
	})
}
