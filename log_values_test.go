package observability

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that Msg returns the message of the log values
func TestLogValues_Msg(t *testing.T) {
	lv := LogValues[string]{msg: "hello"}
	assert.Equal(t, "hello", lv.Msg())
}

// Test that Err returns the error of the log values
func TestLogValues_Err(t *testing.T) {
	err := errors.New("oops")
	lv := LogValues[string]{err: err}
	assert.Equal(t, err, lv.Err())
}

// Test that DebugValues returns the debug values of the log values
func TestLogValues_DebugValues(t *testing.T) {
	lv := LogValues[string]{debugValues: []string{"foo", "bar"}}
	assert.EqualValues(t, []string{"foo", "bar"}, lv.DebugValues())
}

// Test that ErrorValues returns the error values of the log values
func TestLogValues_ErrorValues(t *testing.T) {
	lv := LogValues[string]{errorValues: []string{"baz", "qux"}}
	assert.EqualValues(t, []string{"baz", "qux"}, lv.ErrorValues())
}

// Test that InfoValues returns the info values of the log values
func TestLogValues_InfoValues(t *testing.T) {
	lv := LogValues[string]{infoValues: []string{"corge", "grault"}}
	assert.EqualValues(t, []string{"corge", "grault"}, lv.InfoValues())
}

// Test that WithMsg sets the message of the builder
func TestLogValuesBuilder_WithMsg(t *testing.T) {
	b := NewLogValuesBuilder[string]().WithMsg("hello")
	assert.Equal(t, "hello", b.msg)
}

// Test that WithInfoValue appends an info value to the builder
func TestLogValuesBuilder_WithInfoValue(t *testing.T) {
	b := NewLogValuesBuilder[string]().WithInfoValue("foo")
	assert.Equal(t, []string{"foo"}, b.infoValues)
}

// Test that WithDebugValue appends a debug value to the builder
func TestLogValuesBuilder_WithDebugValue(t *testing.T) {
	b := NewLogValuesBuilder[string]().WithDebugValue("bar")
	assert.Equal(t, []string{"bar"}, b.debugValues)
}

// Test that WithErrorValue appends an error value to the builder
func TestLogValuesBuilder_WithErrorValue(t *testing.T) {
	b := NewLogValuesBuilder[string]().WithErrorValue("baz")
	assert.Equal(t, []string{"baz"}, b.errorValues)
}

// Test that Build creates a log values with the builder's fields
func TestLogValuesBuilder_Build(t *testing.T) {
	b := NewLogValuesBuilder[string]()
	b.msg = "hello"
	b.err = errors.New("oops")
	b.infoValues = []string{"corge"}
	b.debugValues = []string{"grault"}
	b.errorValues = []string{"garply"}

	lv := b.Build()
	assert.Equal(t, "hello", lv.msg)
	assert.Equal(t, errors.New("oops"), lv.err)
	assert.EqualValues(t, []string{"corge"}, lv.infoValues)
	assert.EqualValues(t, []string{"grault"}, lv.debugValues)
	assert.EqualValues(t, []string{"garply"}, lv.errorValues)
}

// Test that NewLogValuesBuilder creates a new empty builder
func TestNewLogValuesBuilder(t *testing.T) {
	b := NewLogValuesBuilder[string]()
	assert.Empty(t, b.msg)
	assert.Nil(t, b.err)
	assert.Empty(t, b.infoValues)
	assert.Empty(t, b.debugValues)
	assert.Empty(t, b.errorValues)
}

// Test that FactorLogValuesBuilder creates a builder with the given options
func TestFactoryLogValuesBuilder(t *testing.T) {
	options := LogValuesOptions[string]{
		{debug: true, Attr: "foo"},
		{err: true, Attr: "bar"},
		{info: true, Attr: "baz"},
	}

	builder := FactoryLogValuesBuilder(options)
	assert.Equal(t, []string{"foo"}, builder.debugValues)
	assert.Equal(t, []string{"bar"}, builder.errorValues)
	assert.Equal(t, []string{"baz"}, builder.infoValues)
}
