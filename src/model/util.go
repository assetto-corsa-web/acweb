package model

import (
	"database/sql"
	"db"
)

func createTxIfRequired(tx *sql.Tx) (*sql.Tx, error) {
	if tx == nil {
		return db.Get().Begin()
	}

	return tx, nil
}
