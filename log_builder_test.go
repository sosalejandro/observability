package observability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that CreateLogValueOption creates a new log value option with the given attribute
func TestLogBuilder_CreateLogValueOption(t *testing.T) {
	lb := &LogBuilder[string]{}
	lvo := lb.CreateLogValueOption("foo")
	assert.Equal(t, "foo", lvo.Attr)
}

// Test that CreateLogValuesOptions creates a new empty log values options
func TestLogBuilder_CreateLogValuesOptions(t *testing.T) {
	lb := &LogBuilder[string]{}
	lvos := lb.CreateLogValuesOptions()
	assert.Empty(t, *lvos)
}

// Test that CreateLogValuesOptionsWith creates a new log values options with the given options
func TestLogBuilder_CreateLogValuesOptionsWith(t *testing.T) {
	lb := &LogBuilder[string]{}
	lvo1 := lb.CreateLogValueOption("foo")
	lvo2 := lb.CreateLogValueOption("bar")
	lvos := lb.CreateLogValuesOptionsWith(lvo1, lvo2)
	assert.Equal(t, LogValuesOptions[string]{lvo1, lvo2}, *lvos)
}

// Test that FactoryLogValuesBuilder creates a builder with the given options
func TestLogBuilder_FactoryLogValuesBuilder(t *testing.T) {
	lb := &LogBuilder[string]{}
	options := LogValuesOptions[string]{
		{debug: true, Attr: "foo"},
		{err: true, Attr: "bar"},
		{info: true, Attr: "baz"},
	}

	builder := lb.FactoryLogValuesBuilder(options)
	assert.Equal(t, []string{"foo"}, builder.debugValues)
	assert.Equal(t, []string{"bar"}, builder.errorValues)
	assert.Equal(t, []string{"baz"}, builder.infoValues)
}

// Test that NewLogBuilder creates a new LogBuilder with the given type
func TestLogBuilder_NewLogBuilder(t *testing.T) {
	lb := NewLogBuilder[string]()
	assert.NotNil(t, lb)
}
