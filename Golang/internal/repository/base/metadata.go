package base

import "time"

type Metadata struct {
	CreatedBy int       `json:"created_by" db:"created_by"`
	UpdatedBy int       `json:"updated_by" db:"updated_by"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ExtraAttribute struct {
	IsActive bool `json:"is_active" db:"is_active"`
}
