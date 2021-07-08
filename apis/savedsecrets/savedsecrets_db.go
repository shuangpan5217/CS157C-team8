package savedsecrets

import "time"

var SavedSecretsTableName = "saved_secret"

type SavedSecrets struct {
	SecretID    string     `json:"secret_id"`
	Username    string     `json:"username"`
	Content     string     `json:"content"`
	Nickname    string     `json:"nickname"`
	CreatedTime *time.Time `json:"created_time"`
}
