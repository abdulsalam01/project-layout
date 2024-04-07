package database

import (
	"github.com/api-sekejap/internal/constant"
	"github.com/jackc/pgx/v5/pgconn"
)

func WrapDuplicateKeyValueErr(err error) error {
	// Attempt to assert the error to *pq.Error.
	if pqErr, ok := err.(*pgconn.PgError); ok {
		if pqErr.Code == constant.ErrDuplicateKeyValue {
			return nil // Skip if it's a unique constraint violation.
		}
	}

	return err
}
