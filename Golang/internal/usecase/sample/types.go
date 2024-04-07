package sample

import (
	"context"

	"github.com/api-sekejap/internal/entity"
)

type sampleResource interface {
	Create(ctx context.Context, params entity.Sample) (int, error)
}

type sampleUsecase struct {
	sampleRepo sampleResource
}
