package sample

import (
	"context"

	"github.com/api-sekejap/internal/entity"
)

func (c *sampleUsecase) Create(ctx context.Context, params entity.Sample) error {
	_, err := c.sampleRepo.Create(ctx, params)
	if err != nil {
		return err
	}

	return nil
}
