package observability

// LogBuilder is a builder for log values
type LogBuilder[T any] struct{}

// NewLogBuilder creates a new LogBuilder with the given type
func NewLogBuilder[T any]() *LogBuilder[T] {
	return &LogBuilder[T]{}
}

// CreateLogValueOption creates a new log value option with the given attribute
func (lb *LogBuilder[T]) CreateLogValueOption(attr T) *LogValueOption[T] {
	return NewLogValueOption[T](attr)
}

// CreateLogValuesOptions creates a new empty log values options
func (lb *LogBuilder[T]) CreateLogValuesOptions() *LogValuesOptions[T] {
	return NewLogValuesOptions[T]()
}

// CreateLogValuesOptionsWith creates a new log values options with the given options
func (lb *LogBuilder[T]) CreateLogValuesOptionsWith(options ...*LogValueOption[T]) *LogValuesOptions[T] {
	var lvos LogValuesOptions[T]
	lvos = append(lvos, options...)
	return &lvos
}

// FactoryLogValuesBuilder creates a builder with the given options
func (lb *LogBuilder[T]) FactoryLogValuesBuilder(options LogValuesOptions[T]) *LogValuesBuilder[T] {
	return FactoryLogValuesBuilder[T](options)
}

// CreateLogValuesBuilder creates a new log values builder
func (lb *LogBuilder[T]) CreateLogValuesBuilder() *LogValuesBuilder[T] {
	return NewLogValuesBuilder[T]()
}
