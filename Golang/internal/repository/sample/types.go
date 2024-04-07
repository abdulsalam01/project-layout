package sample

import (
	"context"

	"github.com/api-sekejap/internal/repository/base"
	"github.com/jackc/pgx/v5/pgconn"
)

type Sample struct {
	ID   int    `json:"id" db:"id" table:"channels"`
	Name string `json:"name" db:"name"`

	base.Metadata
	base.ExtraAttribute
}

type databaseResource interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type ChannelsRepo struct {
	database databaseResource
}
