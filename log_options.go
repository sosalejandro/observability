package observability

// LogValueOption is an option for a log value
type LogValueOption[T any] struct {
	Attr             T
	debug, info, err bool
}

// NewLogValueOption creates a new log value option with the given attribute
func NewLogValueOption[T any](attr T) *LogValueOption[T] {
	return &LogValueOption[T]{Attr: attr}
}

// WithDebug sets the debug flag to true
func (o *LogValueOption[T]) WithDebug() *LogValueOption[T] {
	o.debug = true
	return o
}

// WithInfo sets the info flag to true
func (o *LogValueOption[T]) WithInfo() *LogValueOption[T] {
	o.info = true
	return o
}

// WithError sets the error flag to true
func (o *LogValueOption[T]) WithError() *LogValueOption[T] {
	o.err = true
	return o
}

// LogValuesOptions is a list of log value options
type LogValuesOptions[T any] []*LogValueOption[T]

// Len returns the length of the list
func (lvos *LogValuesOptions[T]) Len() int {
	return len(*lvos)
}

// AddOption adds the given option to the list
type LogValueLevelOption struct {
	info, err, debug bool
}

// LogValuesBuilder is a builder for log values
func NewLogValuesOptions[T any]() *LogValuesOptions[T] {
	return &LogValuesOptions[T]{}
}

// AddOption adds the given option to the list
func (lvos *LogValuesOptions[T]) AddOption(option *LogValueOption[T]) *LogValuesOptions[T] {
	*lvos = append(*lvos, option)
	return lvos
}

// NewLogValuesLevelOption creates a new log value level option
func NewLogValuesLevelOption(info, err, debug bool) *LogValueLevelOption {
	return &LogValueLevelOption{
		info:  info,
		err:   err,
		debug: debug,
	}
}
