package tests

import (
	"os"

	dblocal "github.com/re-worthy/backend-go/internal/db/local"
	dbshared "github.com/re-worthy/backend-go/internal/db/shared"
	handlers "github.com/re-worthy/backend-go/internal/handlers/types"
)

func NewTestBaseHandler() (*handlers.TBaseHandler, dbshared.TOnClose, error) {
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, func() error { return nil }, err
	}
	db, onclose, err := dblocal.GetLocalConnection("file:" + dir + "/test.db")
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return &handlers.TBaseHandler{DB: db}, onclose, nil
}
