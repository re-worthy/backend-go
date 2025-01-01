package db_remote

import (
	"database/sql"
	"errors"
	"strings"

	db_shared "github.com/re-worthy/backend-go/internal/db/shared"

	_ "github.com/re-worthy/libsql-client-go/libsql"
)

const DRIVER_NAME = "libsql-old"

var (
	ErrDBPrefix      = errors.New("Expected database connection string to contain valid prefix")
	ErrDBDomainNAuth = errors.New("Expected database connection string to contain valid domain name and auth url param")
)

func validateDbUrl(str string) error {
	hasPrefix := strings.HasPrefix(str, "libsql://")
	if !hasPrefix {
		return ErrDBPrefix
	}
	hasDomainNAuth := strings.Contains(str, ".turso.io?authToken=")
	if !hasDomainNAuth {
		return ErrDBDomainNAuth
	}
	return nil
}

/*
Dont forget to defer close database
*/
func GetRemoteConnection(dataSourceName string) (*sql.DB, db_shared.TOnClose, error) {
	urlError := validateDbUrl(dataSourceName)
	if urlError != nil {
		return &sql.DB{}, func() error { return nil }, urlError
	}

	db, err := sql.Open(DRIVER_NAME, dataSourceName)
	if err != nil {
		return &sql.DB{}, func() error { return nil }, err
	}

	return db, db_shared.CreateOnCloseSuccess(db), nil
}
