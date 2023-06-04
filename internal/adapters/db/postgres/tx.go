package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	errorTxNotStarted = func(err error) error {
		return fmt.Errorf("failed to start transaction %+v", err)
	}
	errorTxCommitted = func(err error) error {
		return fmt.Errorf("failedt to commit transaction: %+v", err)
	}
)

func rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		if err != sql.ErrTxDone {
			log.Printf("failed to rollback transaction: %+v", err)
		}
	}
}
