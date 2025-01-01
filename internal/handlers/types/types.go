package handlers

import (
	"database/sql"
	"net/http"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
)

type ResponseError struct {
	Err         error // treated as an actual error occured in proram
	User_err    error // used in response
	Status_code int
}

type (
	THandlerFunc[Rq any, Rs any] func(r *http.Request, w http.ResponseWriter, reqBodySchema *Rq, generalHandler *TBaseHandler) (*Rs, *ResponseError)
	TBaseHandler                 struct {
		DB      *sql.DB
		Queries *gen.Queries
		/* deps for all routes will be here */
	}
)
