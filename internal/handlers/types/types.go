package handlers

import (
	"database/sql"
	"net/http"
)

type (
	THandlerFunc[Rq any, Rs any] func(r *http.Request, w http.ResponseWriter, reqBodySchema *Rq, generalHandler *TBaseHandler) (error, *Rs)
	TBaseHandler                 struct {
		DB *sql.DB
		/* deps for all routes will be here */
	}
)
