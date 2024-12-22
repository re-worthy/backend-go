package dblocal

import (
	"database/sql"

	db_shared "github.com/re-worthy/backend-go/internal/db/shared"

	_ "github.com/tursodatabase/go-libsql"
)

/*
Dont forget to defer close database
*/
func GetLocalConnection(dataSourceName string) (*sql.DB, db_shared.TOnClose, error) {
	db, err := sql.Open(db_shared.DRIVER_NAME, dataSourceName)
	if err != nil {
		return &sql.DB{}, func() error { return nil }, err
	}

	return db, db_shared.CreateOnCloseSuccess(db), nil
}
