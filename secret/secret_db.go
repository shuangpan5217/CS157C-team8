package secret

import "time"

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
