package tools

import (
	"context"

	db "github.com/api-sekejap/pkg/database"
)

type seederRunner struct {
	Data map[string]interface{} `json:"data"`
	Type map[string]string      `json:"type"`
}
type seederType struct {
	Type string `json:"type"`
}

// Seeder interface for all seeders.
type seederResources interface {
	Seed(ctx context.Context, data seederRunner, base db.DatabaseHelper) error
}

// Implement Seeder for all tables.
type SampleSeeder struct{}
