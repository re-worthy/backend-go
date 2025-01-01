package db_init

import (
	"database/sql"
)

/*
* Should be removed when the App will be deployed
* Now database inits when applying migrations (make goose up)
* TODO remove
* */
func InitDB(db *sql.DB) error {
	var err error
	_, err = db.Exec("CREATE TABLE counter (id INTEGER PRIMARY KEY, cnt INTEGER)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO counter (id, cnt) VALUES (1, 0)")
	if err != nil {
		return err
	}

	return nil
}
