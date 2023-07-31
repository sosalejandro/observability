package observability

type DebugValues[T any] []T
type ErrorValues[T any] []T
type InfoValues[T any] []T

// LogValues is a wrapper for log values to be passed to the logger
type LogValues[T any] struct {
	msg         string
	err         error
	debugValues DebugValues[T]
	errorValues ErrorValues[T]
	infoValues  InfoValues[T]
}

// Msg returns the message
func (lv LogValues[T]) Msg() string {
	return lv.msg
}

// Err returns the error
func (lv LogValues[T]) Err() error {
	return lv.err
}

// DebugValues returns the debug values
func (lv LogValues[T]) DebugValues() DebugValues[T] {
	return lv.debugValues
}

// ErrorValues returns the error values
func (lv LogValues[T]) ErrorValues() ErrorValues[T] {
	return lv.errorValues
}

// InfoValues returns the info values
func (lv LogValues[T]) InfoValues() InfoValues[T] {
	return lv.infoValues
}

// LogValuesBuilder is a builder for log values
type LogValuesBuilder[T any] struct {
	msg         string
	err         error
	infoValues  []T
	debugValues []T
	errorValues []T
}

// WithMsg sets the message
func (b *LogValuesBuilder[T]) WithMsg(msg string) *LogValuesBuilder[T] {
	b.msg = msg
	return b
}

// WithErr sets the error
func (b *LogValuesBuilder[T]) WithInfoValue(field T) *LogValuesBuilder[T] {
	b.infoValues = append(b.infoValues, field)
	return b
}

// WithDebugValue sets the debug value
func (b *LogValuesBuilder[T]) WithDebugValue(field T) *LogValuesBuilder[T] {
	b.debugValues = append(b.debugValues, field)
	return b
}

// WithErrorValue sets the error value
func (b *LogValuesBuilder[T]) WithErrorValue(field T) *LogValuesBuilder[T] {
	b.errorValues = append(b.errorValues, field)
	return b
}

// Build builds the log values
func (b *LogValuesBuilder[T]) Build() LogValues[T] {
	return LogValues[T]{
		msg:         b.msg,
		err:         b.err,
		infoValues:  b.infoValues,
		debugValues: b.debugValues,
		errorValues: b.errorValues,
	}
}

// NewLogValuesBuilder creates a new log values builder
func NewLogValuesBuilder[T any]() *LogValuesBuilder[T] {
	return &LogValuesBuilder[T]{}
}

// FactoryLogValuesBuilder creates a builder with the given options
func FactoryLogValuesBuilder[T any](options LogValuesOptions[T]) *LogValuesBuilder[T] {
	builder := NewLogValuesBuilder[T]()
	for _, option := range options {
		switch {
		case option.debug:
			builder = builder.WithDebugValue(option.Attr)
		case option.err:
			builder = builder.WithErrorValue(option.Attr)
		case option.info:
			builder = builder.WithInfoValue(option.Attr)
		}
	}
	return builder
}
