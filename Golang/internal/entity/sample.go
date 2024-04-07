package entity

import "github.com/api-sekejap/internal/entity/base"

type Sample struct {
	ID   int    `json:"id"`
	Name string `json:"name"`

	base.Metadata
	base.ExtraAttribute
}
