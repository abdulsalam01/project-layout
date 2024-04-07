package app

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// Private struct helper.
// Goes here:
type dbQueryTracer struct {
	log *logrus.Logger
}

func (tracer *dbQueryTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	tracer.log.Infof("Executing command sql: %s with args: %v", data.SQL, data.Args)
	return ctx
}
func (tracer *dbQueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	// Add logic here...
}
