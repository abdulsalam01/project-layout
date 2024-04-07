package database

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Helper-extenders function.
// Override function initializer.
func (b *DatabaseHelper) WithTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	tx, err := b.Database.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.ReadCommitted,
		AccessMode: pgx.ReadWrite,
	})

	if err != nil {
		return err
	}

	// Defer a rollback in case of panic or error.
	defer func() {
		if r := recover(); r != nil || err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Execute the function within the transaction.
	err = fn(tx)
	if err != nil {
		return err
	}

	tx.Commit(ctx)
	return nil
}
