package handlers

import (
	"database/sql"
	"net/http"

	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
)

type (
	THandlerFunc[Rq any, Rs any] func(r *http.Request, w http.ResponseWriter, reqBodySchema *Rq, generalHandler *TBaseHandler) (error, *Rs)
	TBaseHandler                 struct {
		DB      *sql.DB
		Queries *gen.Queries
		/* deps for all routes will be here */
	}
)
