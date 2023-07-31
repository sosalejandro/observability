package observability

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test that NewLogValueOption creates a new option with the given attribute
func TestNewLogValueOption(t *testing.T) {
	option := NewLogValueOption("foo")
	assert.Equal(t, "foo", option.Attr)
	assert.False(t, option.debug)
	assert.False(t, option.info)
	assert.False(t, option.err)
}

// Test that WithDebug sets the debug flag to true
func TestLogValueOption_WithDebug(t *testing.T) {
	option := NewLogValueOption("foo").WithDebug()
	assert.True(t, option.debug)
}

func TestLogValueOption_WithInfo(t *testing.T) { // Test that WithInfo sets the info flag to true
	option := NewLogValueOption("foo").WithInfo()
	assert.True(t, option.info)
}

// Test that WithError sets the err flag to true
func TestLogValueOption_WithError(t *testing.T) {
	option := NewLogValueOption("foo").WithError()
	assert.True(t, option.err)
}

// Test that NewLogValuesOptions creates an empty slice of options
func TestNewLogValuesOptions(t *testing.T) {
	options := NewLogValuesOptions[string]()
	assert.Empty(t, options)
}

// Test that AddOption appends an option to the slice
func TestLogValuesOptions_AddOption(t *testing.T) {
	options := NewLogValuesOptions[string]()
	option := NewLogValueOption("foo")
	options.AddOption(option)
	assert.Equal(t, options.Len(), 1)
	assert.Equal(t, option, (*options)[0])
}

// Test that NewLogValuesLevelOption creates a new option with the given flags
func TestNewLogValuesLevelOption(t *testing.T) {
	option := NewLogValuesLevelOption(true, false, true)
	assert.True(t, option.info)
	assert.False(t, option.err)
	assert.True(t, option.debug)
}
