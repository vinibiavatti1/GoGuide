// Must
// Must is a pattern that allows you to assert that a value is not nil or an error.
// It is often used to simplify error handling when you're sure a function should not fail.
// The Must pattern is typically implemented as a function that takes a value and an error,
// and either returns the value or panics if there's an error.

package patterns

import "fmt"

// Struct
// The struct below will be used to demonstrate the Must pattern.
type Config struct {
	Name  string
	Value string
}

// Constructor
// This constructor returns an error if the name is empty.
// This allows us to validate the input before creating the Config instance.
// However, it makes it inconvenient to use when we're creating safe, static configs.
func NewConfig(name, value string) (*Config, error) {
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	return &Config{Name: name, Value: value}, nil
}

// Must Function
// To avoid checking for errors every time, we can wrap the constructor in a Must function.
// If the constructor fails, this function will panic.
// This is useful when we are confident that the input is valid (e.g., preset configs).
func MustNewConfig(name, value string) *Config {
	config, err := NewConfig(name, value)
	if err != nil {
		panic(err)
	}
	return config
}

// Preset Configurations
// Since we're sure these values are valid, we use the Must function directly.
var Configs = []*Config{
	MustNewConfig("host", "localhost"),
	MustNewConfig("port", "8080"),
}
