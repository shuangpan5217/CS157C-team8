package user

import "time"

var userTableName = "user"

type UserPost struct {
	Username    string     `json:"username" db:"username"`
	Description string     `json:"description"`
	HashValue   string     `json:"hashvalue"`
	Nickname    string     `json:"nickname"`
	Password    string     `json:"password"`
	CreatedTime *time.Time `json:"created_time"`
}

type UserUpdate struct {
	Description string `json:"description"`
	Nickname    string `json:"nickname"`
}

type ErrorResponse struct {
	Message       string
	Error         string
	MessageStatus int
}

type SuccessResponse struct {
	Message       string
	MessageStatus int
}
