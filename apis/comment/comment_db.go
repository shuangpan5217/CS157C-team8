package comment

import "time"

var CommentTableName = "comment"

type CommentPost struct {
	CommentID   string     `json:"comment_id"`
	Comment     string     `json:"comment"`
	CreatedTime *time.Time `json:"created_time"`
	Nickname    string     `json:"nickname"`
	SecretID    string     `json:"secret_id"`
}

type CommentGet struct {
	Comment     string     `json:"comment"`
	CreatedTime *time.Time `json:"created_time"`
	Nickname    string     `json:"nickname"`
}
