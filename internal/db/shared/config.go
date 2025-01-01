package db_shared

import (
	"database/sql"
	"log"
)

type TOnClose = func() error

func CreateOnCloseSuccess(db *sql.DB) TOnClose {
	return func() error {
		if err := db.Close(); err != nil {
			return err
		}
		log.Println("Gracefully handled database.close")
		return nil
	}
}
