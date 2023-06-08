package postgres

import (
	"database/sql"
	"log"

	"github.com/lordvidex/errs"
)

var (
	errorTxNotStarted = func(err error) error {
		return errs.B(err).Code(errs.Unauthenticated).Msg("failed to start transaction").Err()
	}
	errorTxCommitted = func(err error) error {
		return errs.B(err).Code(errs.Unauthenticated).Msg("failedt to commit transaction:").Err()
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
