package zap

import (
	"context"

	"github.com/sosalejandro/observability"
	"go.uber.org/zap"
)

// ZapLogger is a logger that uses Zap
// It doesn't provide any additional functionality over the base ObservabilityLogger.
//
// It doesn't implement context usage for the context methods,
// it requires custom implementation to manipulate the context from your side.
type ZapLogger[T zap.Field] struct {
	logger *zap.Logger
}

// NewZapLogger creates a new ZapLogger with the logger
func NewZapLogger(logger *zap.Logger) observability.ObservabilityLogger[zap.Field] {
	return &ZapLogger[zap.Field]{logger: logger}
}

// LogInfo logs a message at the info level
func (l *ZapLogger[T]) LogInfo(lv observability.LogValues[zap.Field]) {
	l.logger.Info(lv.Msg(), lv.InfoValues()...)
}

// LogDebug logs a message at the debug level
func (l *ZapLogger[T]) LogDebug(lv observability.LogValues[zap.Field]) {
	l.logger.Debug(lv.Msg(), lv.DebugValues()...)
}

// LogError logs a message at the error level
func (l *ZapLogger[T]) LogError(lv observability.LogValues[zap.Field]) {
	l.logger.Error(lv.Msg(), lv.ErrorValues()...)
}

// LogInfoContext logs a message at the info level with a context (requires custom implementation)
func (l *ZapLogger[T]) LogInfoContext(ctx context.Context, lv observability.LogValues[zap.Field]) {
	l.logger.Info(lv.Msg(), lv.InfoValues()...)
}

// LogDebugContext logs a message at the debug level with a context (requires custom implementation)
func (l *ZapLogger[T]) LogDebugContext(ctx context.Context, lv observability.LogValues[zap.Field]) {
	l.logger.Debug(lv.Msg(), lv.DebugValues()...)
}

// LogErrorContext logs a message at the error level with a context (requires custom implementation)
func (l *ZapLogger[T]) LogErrorContext(ctx context.Context, lv observability.LogValues[zap.Field]) {
	l.logger.Error(lv.Msg(), lv.ErrorValues()...)
}
