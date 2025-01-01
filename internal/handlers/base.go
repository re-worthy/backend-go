package handlers

import (
	"database/sql"

	dblocal "github.com/re-worthy/backend-go/internal/db/local"
	dbremote "github.com/re-worthy/backend-go/internal/db/remote"
	dbshared "github.com/re-worthy/backend-go/internal/db/shared"
	gen "github.com/re-worthy/backend-go/internal/db/sqlc/__gen"
	env_init "github.com/re-worthy/backend-go/internal/env"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

func NewBaseHandler(env env_init.TEnvConfig) (*handlers.TBaseHandler, dbshared.TOnClose, error) {
	var db *sql.DB
	var onclose dbshared.TOnClose
	var err error

	if env.ENVIRONMENT == env_init.ENV_NAME_PRODUCTION {
		db, onclose, err = dbremote.GetRemoteConnection(env.DATABASE_URL)
	} else {
		db, onclose, err = dblocal.GetLocalConnection(env.DATABASE_URL)
	}

	if err != nil {
		return nil, func() error { return nil }, err
	}

	q := gen.New(db)

	return &handlers.TBaseHandler{DB: db, Queries: q}, onclose, nil
}
