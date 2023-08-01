package zap

import (
	"context"
	"testing"

	"github.com/sosalejandro/observability"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func setupLogsCapture() (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(zap.DebugLevel)
	return zap.New(core), logs
}

// Create a new ZapLogger with the logger
// Create a LogValues object with some test data
func arrange() (*observer.ObservedLogs, observability.ObservabilityLogger[zap.Field], zapcore.Field, observability.LogValues[zap.Field]) {
	logger, logs := setupLogsCapture()

	zapLogger := NewZapLogger(logger)

	field := zap.String("foo", "bar")

	lv := observability.NewLogValuesBuilder[zap.Field]().
		WithInfoValue(field).
		WithDebugValue(field).
		WithErrorValue(field).
		WithMsg("test message").
		Build()
	return logs, zapLogger, field, lv
}

func TestZapLogger_LogInfo(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogInfo method on the ZapLogger
	zapLogger.LogInfo(lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.InfoLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}

func TestZapLogger_LogDebug(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogDebug method on the ZapLogger
	zapLogger.LogDebug(lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.DebugLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}

func TestZapLogger_LogError(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogError method on the ZapLogger
	zapLogger.LogError(lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.ErrorLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}

func TestZapLogger_LogInfoContext(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogInfo method on the ZapLogger
	zapLogger.LogInfoContext(context.Background(), lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.InfoLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}

func TestZapLogger_LogDebugContext(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogDebug method on the ZapLogger
	zapLogger.LogDebugContext(context.Background(), lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.DebugLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}

func TestZapLogger_LogErrorContext(t *testing.T) {
	logs, zapLogger, field, lv := arrange()

	// Call the LogError method on the ZapLogger
	zapLogger.LogErrorContext(context.Background(), lv)

	// Assert logger has a value
	assert.Equal(t, logs.Len(), 1)
	entry := logs.All()[0]
	// Assert entry has the correct values
	assert.Equal(t, entry.Level, zap.ErrorLevel)
	assert.Equal(t, entry.Message, "test message")
	assert.Equal(t, entry.Context, []zap.Field{field})
}
