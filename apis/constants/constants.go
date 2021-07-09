package constants

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/gocql/gocql"
)

// Session for Cassandra
var Session *gocql.Session

type ErrorResponse struct {
	Message    string
	Error      string
	StatusCode int
}

type SuccessResponse struct {
	Message        string
	SuccessMessage string
	StatusCode     int
	Body           interface{}
}

func GenerateErrorResponse(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	w.WriteHeader(statusCode)
	render.JSON(w, r, ErrorResponse{
		Message:    "Error encountered",
		Error:      err.Error(),
		StatusCode: statusCode,
	})
}
