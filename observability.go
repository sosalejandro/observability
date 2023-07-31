package observability

import (
	"context"
	"errors"
	"reflect"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TracesValues contains the traceId and spanId.
//
// Requires to setup zapcore.ObjectMarshaler  on zap.Object & slog.GroupValue on slog.Group
// transformations to zap.Field & slog.Attr respectively.
// Refer to https://pkg.go.dev/go.uber.org/zap#Object & https://pkg.go.dev/github.com/x/slog#GroupValue
type TraceValues struct {
	TraceId string
	SpanId  string
}

type ObservabilityLogging[T any] interface {
	// CreateLogBuilder creates a new LogBuilder for compatible log values with the given type
	// Receives a function that sets up tracing for the log values
	CreateLogBuilder() *LogBuilder[T]
	// LogInfo logs an info message with the given values
	LogInfo(lv LogValues[T], opts ...trace.EventOption)
	// LogError logs an error message with the given values
	LogError(lv LogValues[T], opts ...trace.EventOption)
	// LogDebug logs a debug message with the given values
	LogDebug(lv LogValues[T], opts ...trace.EventOption)
	// LogInfoContext logs an info message with the given values and observability context
	LogInfoContext(lv LogValues[T], opts ...trace.EventOption)
	// LogErrorContext logs an error message with the given values and observability context
	LogErrorContext(lv LogValues[T], opts ...trace.EventOption)
	// LogDebugContext logs a debug message with the given values and observability context
	LogDebugContext(lv LogValues[T], opts ...trace.EventOption)
}

type ObservabilityLogger[T any] interface {
	// LogInfo logs an info message with the given values
	LogInfo(lv LogValues[T])
	// LogError logs an error message with the given values
	LogError(lv LogValues[T])
	// LogDebug logs a debug message with the given values
	LogDebug(lv LogValues[T])
	// LogInfoContext logs an info message with the given values and observability context
	LogInfoContext(ctx context.Context, lv LogValues[T])
	// LogErrorContext logs an error message with the given values and observability context
	LogErrorContext(ctx context.Context, lv LogValues[T])
	// LogDebugContext logs a debug message with the given values and observability context
	LogDebugContext(ctx context.Context, lv LogValues[T])
}

type ObservabilityHandler[T any] interface {
	// StartSpan starts a span with the given name and options
	// and returns the context and a function to shutdown the span
	StartSpan(name string, opts ...trace.SpanStartOption) (ctx context.Context, shutdown func(...trace.SpanEndOption))
	GetTraceValues() (TraceValues, error)
	// SetTracingFormat sets the tracing format for the given tracingSetup function
	// Requires the SetTracingFormat to not have been called before
	SetTracingFormat(tracingSetup func(string, TraceValues) T) error
	ObservabilityLogging[T]
}

// ObservabilityContext contains the context, span, logger and tracing format
// used for observability
//
// Provides a wrapper for the observability functions triggered with log calls
type ObservabilityContext[T any] struct {
	// serviceName is the name of the service
	serviceName string
	// The context used for observability
	ctx context.Context
	// The span used for observability
	span trace.Span
	// Logger used for observability
	logger ObservabilityLogger[T]
	// traceId is the id of the trace
	traceId string
	// spanId is the id of the span
	spanId string
	// tracingFormat is a zapcore.Field or slog.Attr that contains the traceId and spanId
	// requires setup of zapcore.ObjectEncoder & slog.Group transformations to zapcore.Field & slog.Attr respectively
	// refer to https://pkg.go.dev/golang.org/x/exp/slog#Group and https://github.com/uber-go/zap/blob/v1.24.0/field.go#L399
	tracingFormat T
	// traceOptions are the options used for tracing
	traceOptions []trace.EventOption
}

func NewObservabilityHandler[T any](ctx context.Context, serviceName string, logger ObservabilityLogger[T]) ObservabilityHandler[T] {
	return &ObservabilityContext[T]{
		ctx:         ctx,
		serviceName: serviceName,
		logger:      logger,
	}
}

// StartSpan starts a span with the given name and options
// and returns the context and a function to shutdown the span
func (oc *ObservabilityContext[T]) StartSpan(name string, opts ...trace.SpanStartOption) (context.Context, func(...trace.SpanEndOption)) {
	oc.ctx, oc.span = trace.SpanFromContext(oc.ctx).TracerProvider().Tracer(oc.serviceName).
		Start(oc.ctx, name, opts...)

	oc.traceId = oc.span.SpanContext().TraceID().String()
	oc.spanId = oc.span.SpanContext().SpanID().String()

	oc.traceOptions = make([]trace.EventOption, 0)
	oc.traceOptions = append(oc.traceOptions, trace.WithAttributes(
		attribute.String("traceId", oc.traceId),
		attribute.String("spanId", oc.spanId),
	))

	return oc.ctx, oc.span.End
}

// CreateLogBuilder creates a new LogBuilder for compatible log values with the given type
func (oc *ObservabilityContext[T]) CreateLogBuilder() *LogBuilder[T] {

	lb := NewLogBuilder[T]()
	if oc.span != nil && !isNil[T](oc.tracingFormat) {
		lb.CreateLogValuesBuilder().
			WithDebugValue(oc.tracingFormat).
			WithErrorValue(oc.tracingFormat).
			WithInfoValue(oc.tracingFormat)
	}

	return lb
}

// LogInfo logs an info message with the given values
func (oc *ObservabilityContext[T]) LogInfo(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.AddEvent(
		lv.Msg(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)

	oc.logger.LogInfo(lv)
}

// LogError logs an error message with the given values
func (oc *ObservabilityContext[T]) LogError(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.RecordError(
		lv.Err(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)
	oc.logger.LogError(lv)
}

// LogDebug logs a debug message with the given values
func (oc *ObservabilityContext[T]) LogDebug(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.AddEvent(
		lv.Msg(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)
	oc.logger.LogDebug(lv)
}

// LogInfoContext logs an info message with the given values and observability context
func (oc *ObservabilityContext[T]) LogInfoContext(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.AddEvent(
		lv.Msg(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)
	oc.logger.LogInfoContext(oc.ctx, lv)
}

// LogErrorContext logs an error message with the given values and observability context
func (oc *ObservabilityContext[T]) LogErrorContext(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.RecordError(
		lv.Err(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)
	oc.logger.LogErrorContext(oc.ctx, lv)
}

// LogDebugContext logs a debug message with the given values and observability context
func (oc *ObservabilityContext[T]) LogDebugContext(lv LogValues[T], opts ...trace.EventOption) {
	oc.span.AddEvent(
		lv.Msg(),
		withTraceOptions(oc.traceOptions, opts...)...,
	)
	oc.logger.LogDebugContext(oc.ctx, lv)
}

// GetTraceValues returns the trace values
// Requires the span to have been started
// Returns an error if the span haven't been started yet
func (oc *ObservabilityContext[T]) GetTraceValues() (TraceValues, error) {
	if oc.span == nil {
		return TraceValues{}, errors.New("span haven't been started yet")
	}

	return TraceValues{
		TraceId: oc.traceId,
		SpanId:  oc.spanId,
	}, nil
}

// SetTracingFormat sets the tracing format for the given tracingSetup function
// Requires the SetTracingFormat to not have been called before
// tracingSetup is a function that receives the name of the field and the trace values
// and returns the tracing format
func (oc *ObservabilityContext[T]) SetTracingFormat(tracingSetup func(string, TraceValues) T) error {
	if !isNil[T](oc.tracingFormat) {
		return errors.New("tracing format already set")
	}

	traceValues, _ := oc.GetTraceValues()

	oc.tracingFormat = tracingSetup("tracing", traceValues)

	return nil
}

// isNil checks if the given value is nil
func isNil[T any](value T) bool {
	// Get the value's underlying type
	valueType := reflect.TypeOf(value)

	// Check if the value is the zero value for its type
	return reflect.ValueOf(value).IsZero() && valueType.Kind() != reflect.Ptr
}

// WithTraceOptions adds the given options to the base options
func withTraceOptions(baseOpts []trace.EventOption, opts ...trace.EventOption) []trace.EventOption {
	baseOpts = append(baseOpts, opts...)
	return baseOpts
}
