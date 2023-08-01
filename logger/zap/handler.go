package zap

import (
	"context"

	"github.com/sosalejandro/observability"
	"go.uber.org/zap"
)

func NewZapHandler(ctx context.Context, serviceName string, zapLogger *zap.Logger) observability.ObservabilityHandler[zap.Field] {
	logger := NewZapLogger(zapLogger)
	return observability.NewObservabilityHandler[zap.Field](ctx, serviceName, logger)
}
