package savedsecrets

import (
	"CS157C-TEAM8/apis/constants"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/spf13/cast"
)

const SavedSecretsTableName = "saved_secret"

type SavedSecretPost struct {
	SecretID    string    `json:"secret_id"`
	SecretOwner string    `json:"secret_owner"` // the username who created the secret
	Username    string    `json:"username"`     // the username who saved the secret, added to favorite list
	Content     string    `json:"content"`
	Nickname    string    `json:"nickname"`
	CreatedTime time.Time `json:"created_time"`
}

type SavedSecretRemove struct {
	SecretID *string `json:"secret_id"`
	Username *string `json:"username"`
}

func GeneratePostSavedSecretSuccessResponse(w http.ResponseWriter, r *http.Request, successMessage string, statusCode int, savedSecret SavedSecretPost) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, constants.SuccessResponse{
		Message:        "success",
		SuccessMessage: successMessage,
		StatusCode:     statusCode,
		Body: map[string]string{
			"secret_id":    savedSecret.SecretID,
			"secret_owner": savedSecret.SecretOwner,
			"nickname":     savedSecret.Nickname,
			"username":     savedSecret.Username,
			"created_time": cast.ToString(savedSecret.CreatedTime),
			"content":      savedSecret.Content,
		},
	})
}

func GenerateGetSavedSecretSuccessResponse(w http.ResponseWriter, r *http.Request, successMessage string, statusCode int, savedSecrets []SavedSecretPost) {
	body := make(map[string][]SavedSecretPost)
	body["saved_secrets"] = savedSecrets

	w.WriteHeader(statusCode)
	render.JSON(w, r, constants.SuccessResponse{
		Message:        "success",
		SuccessMessage: successMessage,
		StatusCode:     statusCode,
		Body:           body,
	})
}
